package domainerror

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type ErrorInterceptor struct{}

func NewErrorInterceptor() *ErrorInterceptor {
	return &ErrorInterceptor{}
}

// UnaryServerInterceptor returns a new unary server interceptor that converts errors to gRPC status errors.
func (i *ErrorInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		if err != nil {
			class, _, message := parse(err)
			return nil, status.Error(gRPCStatus(class), message)
		}
		return resp, nil
	}
}
