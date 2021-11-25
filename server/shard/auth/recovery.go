package auth

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UnaryServerInterceptor returns a new unary server interceptor for panic recovery.
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {

		defer func() {
			if r := recover(); r != nil  {
				//err = recoverFrom(ctx, r, o.recoveryHandlerFunc)
				//return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("token not valid: %v", err))
				err = status.Error(codes.Unknown, fmt.Sprintf("panic errr: %v", err))
			}
		}()

		resp, err := handler(ctx, req)
		return resp, err
	}
}
