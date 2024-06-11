package main

import (
	"context"
	"errors"
	"github.com/zhangpetergo/GoLab/slog-example/foundation/logger"
	"math/rand"
	"os"
	"runtime"
)

func main() {
	var log *logger.Logger

	events := logger.Events{
		Error: func(ctx context.Context, r logger.Record) {
			log.Info(ctx, "******* SEND ALERT *******")
		},
	}

	traceIDFn := func(ctx context.Context) string {
		return ""
	}

	log = logger.NewWithEvents(os.Stdout, logger.LevelInfo, "SALES", traceIDFn, events)

	// -------------------------------------------------------------------------

	ctx := context.Background()
	if err := run(ctx, log); err != nil {
		log.Error(ctx, "startup", "message", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *logger.Logger) error {

	log.Info(ctx, "startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))
	defer log.Info(ctx, "shutdown", "GOMAXPROCS", runtime.GOMAXPROCS(0))
	if n := rand.Intn(100) % 2; n == 0 {
		return errors.New("ohh bad thing")
	}

	return nil
}
