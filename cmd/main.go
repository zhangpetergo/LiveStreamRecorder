package main

import (
	"github.com/zhangpetergo/LiveStreamRecorder/foundation/logger"
	"os"
)

/*
 TODO: 1. 使用 ffmpeg 下载直播流
	   2. 打印当前录制的直播流的信息
	   3. 读取 url 的配置文件，支持热更新
	   4. 支持多个 url 同时录制
	   5. 支持对录制的配置

*/

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
