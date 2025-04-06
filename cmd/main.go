package main

import (
	"fmt"
	"github.com/zhangpetergo/LiveStreamRecorder/app/config"
	"github.com/zhangpetergo/LiveStreamRecorder/app/monitor"

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

  下一步：读取 url 配置文件，测试主程序
		优化测试代码

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

	// 加载配置文件
	config.SetConfigPath("./config/config.yaml")
	// 这里主动初始化
	cfg, err := config.GetConfig()
	if err != nil {
		return err
	}

	//url := "https://live.douyin.com/788699151429"
	//url := "https://live.douyin.com/7032984711"

	// 读取 url 配置文件

	// 返回 url 列表
	urls := []string{"https://live.douyin.com/1", "https://live.douyin.com/2", "https://live.douyin.com/3"}
	//names := []string{"1", "2", "3"}

	// 监测直播状态
	monitor.Listen(urls)

	// 创建 Ticker，每 PollIntervalSeconds 秒检查一次直播状态
	checkTicker := time.NewTicker(time.Duration(cfg.PollIntervalSeconds) * time.Second)
	defer checkTicker.Stop()

	// 创建 Ticker，每 5 秒刷新一次控制台输出状态
	printTicker := time.NewTicker(5 * time.Second)
	defer printTicker.Stop()

	// 监听系统信号（Ctrl+C 退出）
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-checkTicker.C:
			// 检查直播状态
			monitor.Listen(urls)
		case <-printTicker.C:
			task.PrintTasks()
		case <-sigChan:
			// 收到退出信号
			fmt.Println("exit")
			return nil
		}

	}
}
