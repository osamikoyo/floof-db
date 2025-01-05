package main

import (
	"context"
	"github.com/osamikoyo/floof-db/internal/app"
	"github.com/osamikoyo/floof-db/pkg/loger"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	loger.New().Error().Err(app.App(os.Args, ctx, loger.New()))
}
