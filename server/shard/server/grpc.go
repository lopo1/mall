package server

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
)


// GRPCConfig defines a grpc server.
type GRPCConfig struct {
	Name              string
	Addr              string
	AuthPublicKeyFile string
	RegisterFunc      func(*grpc.Server)
	Logger            *zap.Logger
}

// RunGRPCServer runs a grpc server.
func RunGRPCServer(c *GRPCConfig) error {
	nameField := zap.String("name", c.Name)
	zap.S().Info("listen addr",c.Addr)
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		zap.S().Fatal("cannot listen", nameField, zap.Error(err))
	}

	//var opts []grpc.ServerOption
	//s := grpc.NewServer(opts...)
	var opts []grpc.UnaryServerInterceptor
	//ts:= auth.GrpcCheckAuth(c.AuthPublicKeyFile)
	//opts = append(opts, ts)
	//role := middlewares.GrpcCheckAdmin()
	//opts = append(opts, role)
	//grpcRecover:=auth.UnaryServerInterceptor()
	//opts = append(opts, grpcRecover)

	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(opts...)))
	c.RegisterFunc(s)
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	zap.S().Info("server started", nameField, zap.String("addr", c.Addr))
	return s.Serve(lis)
}
