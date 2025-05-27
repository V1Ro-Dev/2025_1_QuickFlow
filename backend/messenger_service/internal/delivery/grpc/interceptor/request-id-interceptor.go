package interceptor

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"quickflow/shared/logger"
)

func RequestIDUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	requestID := md.Get(string(logger.RequestID))

	if len(requestID) == 0 {
		// fallback: создаём новый
		requestID = []string{uuid.New().String()}
	}

	ctx = context.WithValue(ctx, logger.RequestID, requestID[0])
	return handler(ctx, req)
}
