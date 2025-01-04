package server

import (
	"github.com/osamikoyo/floof-db/pkg/loger"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	Lis  net.Listener
	Grpc *grpc.Server
	loger.Logger
}

func NewServer() Server {
	lis, err := net.Listen()

	return Server{
		Grpc: grpc.NewServer(),
	}
}

func (s *Server) Run() error {

}
