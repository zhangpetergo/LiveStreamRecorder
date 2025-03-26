package main

import (
	"fmt"
	"github.com/zhangpetergo/LiveStreamRecorder/app/monitor"
	"github.com/zhangpetergo/LiveStreamRecorder/app/processor"
	"github.com/zhangpetergo/LiveStreamRecorder/app/recorder"
	"github.com/zhangpetergo/LiveStreamRecorder/app/resolver/douyin"
	"github.com/zhangpetergo/LiveStreamRecorder/app/task"
	"github.com/zhangpetergo/LiveStreamRecorder/foundation/logger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
 TODO: [x]1. 使用 ffmpeg 下载直播流
	   2. 打印当前录制的直播流的信息
	   3. 读取 url 的配置文件，支持热更新
	   4. 支持多个 url 同时录制
	   5. 支持对录制的配置

  下一步：完成直播状态检测
		测试多个 url 录制的功能

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
	//url := "https://live.douyin.com/788699151429"
	//url := "https://live.douyin.com/7032984711"

	// 读取 url 配置文件

	// 返回 url 列表
	urls := []string{"https://live.douyin.com/1", "https://live.douyin.com/2", "https://live.douyin.com/3"}
	names := []string{"1", "2", "3"}

	// 检测直播状态
	monitor.MonitorStreams(urls)

	// 创建 Ticker，每 60 秒检查一次直播状态
	checkTicker := time.NewTicker(60 * time.Second)
	defer checkTicker.Stop()

	// 创建 Ticker，每 5 秒刷新一次状态
	printTicker := time.NewTicker(5 * time.Second)
	defer printTicker.Stop()

	// 监听系统信号（Ctrl+C 退出）
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-checkTicker.C:
			// 检查直播状态
			monitor.MonitorStreams(urls)
		case <-printTicker.C:
			task.PrintTasks()
		case <-sigChan:
			// 收到退出信号
			fmt.Println("exit")
			return nil
		}

	}
}

func test() {
	url := "https://live.douyin.com/695496496290"
	data, err := douyin.GetStreamData(url)
	if err != nil {
		logger.Log.Error(err.Error())
	}
	err = recorder.Record(data)
	if err != nil {
		logger.Log.Error(err.Error())
	}

}
