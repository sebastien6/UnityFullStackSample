package protocol

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	apiv1 "github.com/sebastien6/UnityFullStackSample/app/pkg/api/v1"
	"github.com/sebastien6/UnityFullStackSample/app/pkg/logger"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func RunGrpcServer(ctx context.Context, v1API apiv1.GamesServiceServer, port string) error {
	// init gRPC listener
	listen, err := net.Listen("tcp", port)
	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("failed to listen: %v", err),
		)
	}
	logger.Log.Sugar().Infof("start listener on port : %s", port)

	opts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
	}

	// gRPC server statup options
	// opts := []grpc.ServerOption{}

	// add middleware
	// opts = middleware.AddLogging(logger.Log, opts)

	// register service
	s := grpc.NewServer(grpc.StreamInterceptor(
		grpc_middleware.ChainStreamServer(
			grpc_recovery.StreamServerInterceptor(),
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger.Log, opts...),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger.Log, opts...),
		)))

	reflection.Register(s)
	apiv1.RegisterGamesServiceServer(s, v1API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			logger.Log.Sugar().Info("shutting down gRPC server...")

			s.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	logger.Log.Sugar().Info("starting gRPC server...")
	return s.Serve(listen)
}
