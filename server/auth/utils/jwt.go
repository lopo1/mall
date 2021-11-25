package utils

import (
	"crypto/rsa"
	"github.com/lopo1/mall/auth/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	ID          uint
	NickName    string
	AuthorityId uint
	jwt.StandardClaims
}

// JWTTokenGen generates a JWT token.
type JWTTokenGen struct {
	privateKey *rsa.PrivateKey
	issuer     string
	nowFunc    func() time.Time
}

// NewJWTTokenGen creates a JWTTokenGen.
func NewJWTTokenGen(issuer string, privateKey *rsa.PrivateKey) *JWTTokenGen {
	return &JWTTokenGen{
		issuer:     issuer,
		nowFunc:    time.Now,
		privateKey: privateKey,
	}
}

// GenerateToken generates a token.
func (t *JWTTokenGen) GenerateToken(user model.User, expire time.Duration) (string, error) {
	nowSec := t.nowFunc().Unix()
	//tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
	//	Issuer:    t.issuer,
	//	IssuedAt:  nowSec,
	//	ExpiresAt: nowSec + int64(expire.Seconds()),
	//	Subject:   accountID,
	//})

	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, CustomClaims{
		ID:             uint(user.ID),
		NickName:       user.NickName,
		AuthorityId:    uint(user.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: nowSec, //签名的生效时间
			ExpiresAt: nowSec + int64(expire.Seconds()),
			Issuer: t.issuer,
		},
	})

	return tkn.SignedString(t.privateKey)
}
