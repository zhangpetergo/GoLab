package main

import (
	"log/slog"
	"os"
)

func main() {

	// 设置 level ，表示设置日志记录的最小等级
	// go 的日志等级从小到大排列
	// debug < info > warning < error

	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)

	logger := slog.New(handler)
	logger.Debug("Debug message")
	logger.Info("Info message")
	logger.Warn("Warning message")
	logger.Error("Error message")

}
