package grpc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/d1zero/scratch/pkg/config"
	e "github.com/d1zero/scratch/pkg/error"
	"go.uber.org/zap"
	gogrpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
)

type Config struct {
	Host string `koanf:"host" validate:"required"`
	Port string `koanf:"port" validate:"required"`
}

type Server struct {
	grpcServer *gogrpc.Server
	logger     *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger, errCodeMap map[e.ErrCode]codes.Code) *Server {
	opts := []gogrpc.ServerOption{
		gogrpc.ChainUnaryInterceptor(
			newErrorUnaryInterceptor(logger, errCodeMap),
		),
	}

	grpcServer := gogrpc.NewServer(opts...)
	reflection.Register(grpcServer)

	return &Server{
		grpcServer: grpcServer,
		logger:     logger,
	}
}

// Server returns gRPC server instance
func (s *Server) Server() *gogrpc.Server {
	return s.grpcServer
}

// Start runs gRPC server with health check
func (s *Server) Start(grpcConfig config.HostPort, healthConfig config.HostPort) {
	go func() {
		addr := net.JoinHostPort(grpcConfig.Host, grpcConfig.Port)
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			s.logger.Errorf("error on tcp socket: %s", err)
			return
		}
		defer lis.Close()

		s.logger.Infof("starting GRPC server on %s", net.JoinHostPort(grpcConfig.Host, grpcConfig.Port))

		err = s.grpcServer.Serve(lis)
		if err != nil {
			s.logger.Errorf("error while serving GRPC server: %s", err)
			return
		}
	}()
}

// Stop gracefully shuts down gRPC server
func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
	s.logger.Info("grpc server shutted down successfully")
}

func newErrorUnaryInterceptor(logger *zap.SugaredLogger, errCodeMap map[e.ErrCode]codes.Code) func(
	ctx context.Context,
	req interface{},
	info *gogrpc.UnaryServerInfo,
	handler gogrpc.UnaryHandler,
) (resp interface{}, err error) {
	return func(
		ctx context.Context,
		req interface{},
		info *gogrpc.UnaryServerInfo,
		handler gogrpc.UnaryHandler,
	) (resp interface{}, err error) {
		d, _ := json.Marshal(req)

		logger.Infof("request params: %s", string(d))

		res, err := handler(ctx, req)
		if err == nil {
			d, _ = json.Marshal(res)
			logger.Infof("response data: %s", string(d))
			return res, err
		}

		logger.Infof("error: %s", err)

		var appErr *e.Error
		if errors.As(err, &appErr) {
			c, ok := errCodeMap[appErr.Code()]
			if !ok {
				c = codes.Unknown
			}

			return nil, status.Error(c, err.Error())
		}

		return nil, status.Error(codes.Unknown, err.Error())
	}
}
