package app

import "github.com/osamikoyo/floof-db/internal/server"

func App() error {
	s := server.NewServer()
	return s.Run()
}
