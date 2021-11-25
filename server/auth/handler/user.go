package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lopo1/mall/auth/global"
	"github.com/lopo1/mall/auth/model"
	"github.com/lopo1/mall/auth/proto"
	"github.com/lopo1/mall/auth/utils"
	"go.uber.org/zap"

	//"github.com/lopo1/mall/auth/utils"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserServer struct{
	TokenExpire    time.Duration
	TokenGenerator TokenGenerator
}
// TokenGenerator generates a token for the specified account.
type TokenGenerator interface {
	GenerateToken(user model.User, expire time.Duration) (string, error)
}
func ModelToRsponse(user model.User) proto.UserInfoResponse{
	//在grpc的message中字段有默认值，你不能随便赋值nil进去，容易出错
	//这里要搞清， 哪些字段是有默认值
	userInfoRsp := proto.UserInfoResponse{
		Id:       user.ID,
		PassWord: user.Password,
		NickName: user.NickName,
		Gender: user.Gender,
		Role: int32(user.Role),
		Mobile: user.Mobile,
	}
	if user.Birthday != nil {
		userInfoRsp.BirthDay = uint64(user.Birthday.Unix())
	}
	return userInfoRsp
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (s *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error){
	//获取用户列表
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println("用户列表")
	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)

	global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)

	for _, user := range users{
		userInfoRsp := ModelToRsponse(user)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}
	return rsp, nil
}

func (s *UserServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error){
	//通过手机号码查询用户
	var user model.User
	result := global.DB.Where(&model.User{Mobile:req.Mobile}).First(&user)
	fmt.Println("GetUserByMobile Mobile=",req.Mobile)
	fmt.Println("GetUserByMobile user=",user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	userInfoRsp := ModelToRsponse(user)
	return &userInfoRsp, nil
}

func (s *UserServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error){
	//通过id查询用户
	var user model.User
	fmt.Println("id=",req.Id)
	result := global.DB.First(&user, req.Id)
	fmt.Println("user=",user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	userInfoRsp := ModelToRsponse(user)
	return &userInfoRsp, nil
}

func (s *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error){
	//新建用户
	var user model.User
	result := global.DB.Where(&model.User{Mobile:req.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	user.Mobile = req.Mobile
	user.NickName = req.NickName

	//密码加密
	//options := &password.Options{16, 100, 32, sha512.New}
	//salt, encodedPwd := password.Encode(req.PassWord, options)
	//user.Password = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	user.Password = utils.HashAndSalt([]byte(req.PassWord))

	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	userInfoRsp := ModelToRsponse(user)
	return &userInfoRsp, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*empty.Empty, error){
	//个人中心更新用户
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	birthDay := time.Unix(int64(req.BirthDay), 0)
	user.NickName = req.NickName
	user.Birthday = &birthDay
	user.Gender = req.Gender

	result = global.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil
}

func (s *UserServer) CheckPassWord(ctx context.Context, req *proto.PasswordCheckInfo) (*proto.CheckResponse, error){
	//校验密码
	options := &password.Options{16, 100, 32, sha512.New}
	passwordInfo := strings.Split(req.EncryptedPassword, "$")
	check := password.Verify(req.Password, passwordInfo[2], passwordInfo[3], options)
	//utils.ComparePasswords()
	return &proto.CheckResponse{Success: check}, nil
}

// Login return token
func (s *UserServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error){
	var user model.User
	result := global.DB.Where("mobile=?",req.Account).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	//校验密码
	//options := &password.Options{16, 100, 32, sha512.New}
	//passwordInfo := strings.Split(user.Password, "$")
	//check := password.Verify(req.Password, passwordInfo[2], passwordInfo[3], options)
	check := utils.ComparePasswords(user.Password,[]byte(req.Password))
	if !check{
		return  nil, status.Error(codes.Internal, "")
	}
	tkn, err := s.TokenGenerator.GenerateToken(user, s.TokenExpire)
	if err != nil {
		zap.S().Error("cannot generate token", zap.Error(err))
		return nil, status.Error(codes.Internal, "")
	}
	return &proto.LoginResponse{AccessToken: tkn,ExpiresIn: int32(s.TokenExpire)}, nil
}