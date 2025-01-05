package app

import (
	"context"
	"github.com/osamikoyo/floof-db/internal/server"
	"github.com/osamikoyo/floof-db/pkg/loger"
)

func App(args []string, ctx context.Context, logger loger.Logger) error {
	s := server.NewServer()

	go func() {
		<-ctx.Done()
		logger.Info().Msg("Floof server stoped! :3")
		s.Shutdown()
	}()

	switch args[1] {
	case "start":
		return s.Run()
	case "stop":
		s.Shutdown()
		return nil
	}

}
