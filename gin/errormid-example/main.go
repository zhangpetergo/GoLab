package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zhangpetergo/GoLab/gin/errormid-example/errs"
	"github.com/zhangpetergo/GoLab/gin/errormid-example/mid"
	"github.com/zhangpetergo/GoLab/slog-example/foundation/logger"
	"math/rand"
	"os"
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

	r := gin.Default()

	r.Use(mid.Error(log))
	r.GET("/testerror", testerror)

	r.Run("0.0.0.0:9999")
}

func testerror(c *gin.Context) {
	if n := rand.Intn(100); n%2 == 0 {
		c.Error(errs.Newf(errs.FailedPrecondition, "this message is trused"))
		return
	}

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	c.JSON(200, status)
}
