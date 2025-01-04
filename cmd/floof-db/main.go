package main

import (
	"github.com/osamikoyo/floof-db/internal/app"
	"github.com/osamikoyo/floof-db/pkg/loger"
)

func main() {
	loger.New().Error().Err(app.App())
}
