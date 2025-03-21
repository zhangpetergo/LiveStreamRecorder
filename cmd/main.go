package main

import (
	"github.com/zhangpetergo/LiveStreamRecorder/foundation/logger"
	"os"
)

func main() {
	// 初始化日志
	logger.InitLogger()

	logger.Log.Info("start")

	if err := run(); err != nil {
		logger.Log.Errorw("startup", "err", err)
		os.Exit(1)
	}
}

func run() error {
	return nil
}
