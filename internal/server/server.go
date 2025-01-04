package server

import (
	"fmt"
	"github.com/osamikoyo/floof-db/internal/config"
	"github.com/osamikoyo/floof-db/pkg/loger"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	Lis    net.Listener
	Grpc   *grpc.Server
	logger loger.Logger
	cfg    config.Config
}

func NewServer() Server {
	cfg := config.Get()

	lis, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		loger.New().Error().Err(err)
		return Server{}
	}

	return Server{
		Grpc:   grpc.NewServer(),
		Lis:    lis,
		logger: loger.New(),
		cfg:    cfg,
	}
}

func (s *Server) Run() error {
	s.logger.Info().Msg(fmt.Sprintf("Floofy-Grpc server started on %s", s.cfg.Addr))

	return s.Grpc.Serve(s.Lis)
}
