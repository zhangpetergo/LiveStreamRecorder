package processor

import (
	"fmt"
	"github.com/zhangpetergo/LiveStreamRecorder/app/recorder"
	"github.com/zhangpetergo/LiveStreamRecorder/app/resolver/douyin"
	"github.com/zhangpetergo/LiveStreamRecorder/app/task"
	"github.com/zhangpetergo/LiveStreamRecorder/foundation/logger"
	"time"
)

// ProcessStream 获取流数据并记录
func ProcessStream(url string) error {
	// 获取流数据
	data, err := douyin.GetStreamData(url)
	if err != nil {
		logger.Log.Errorw("GetStreamData", "url", url, "err", err)
		return err
	}

	// 记录流数据
	err = recorder.Record(data)
	if err != nil {
		logger.Log.Errorw("GetStreamData", "url", url, "err", err)
		return err
	}

	return nil
}

func MockProcessStream(url string, name string) error {

	// start 之后 添加 任务 到 task 数组中
	start(url)

	task.AddTask(url, name)

	// 如果任务结束或者发生任务错误，从 task 数组中删除任务

	wait(url, name)
	task.RemoveTask(url)
	return nil
}

func start(url string) {
	fmt.Println("开始录制直播...", url)
}

func wait(url string, name string) {
	var sleep int

	switch name {
	case "1":
		sleep = 10
	case "2":
		sleep = 20
	case "3":
		sleep = 30
	}

	done := make(chan bool)
	go func() {

		time.Sleep(time.Duration(sleep) * time.Second)
		done <- true
	}()
	select {
	case <-done:
		fmt.Println("录制完成", url)
	}
}
