package task

import (
	"fmt"
	"sync"
	"time"
)

// RecordingTask 结构体，表示一个直播录制任务
type RecordingTask struct {
	StreamURL string
	Name      string
}

// 任务管理器（全局变量）
var (
	tasks []RecordingTask
	mu    sync.RWMutex
)

// AddTask 添加任务
func AddTask(streamURL, name string) {
	mu.Lock()
	defer mu.Unlock()

	tasks = append(tasks, RecordingTask{
		StreamURL: streamURL,
		Name:      name,
	})
}

// GetTasks 获取所有任务
func GetTasks() []RecordingTask {
	mu.RLock()
	defer mu.RUnlock()

	// 返回任务列表的副本，避免外部修改
	tasksCopy := make([]RecordingTask, len(tasks))
	copy(tasksCopy, tasks)
	return tasksCopy
}

// RemoveTask 移除任务
func RemoveTask(streamURL string) {
	mu.Lock()
	defer mu.Unlock()

	for i, task := range tasks {
		if task.StreamURL == streamURL {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return
		}
	}
}

// PrintTasks 打印当前任务
func PrintTasks() {
	// ANSI 代码：\033[H 将光标移到左上角，\033[J 清屏
	fmt.Print("\033[H\033[J")                           // 清空终端
	fmt.Println("\033[1;32m当前正在录制的直播（每 5 秒刷新）：\033[0m") // 绿色标题

	tasks = GetTasks()
	for _, task := range tasks {
		// 打印当前时间
		fmt.Println(task.Name, "正在录制", time.Now().Format("2006-01-02 15:04:05"))
	}
	// 确保刷新立即生效
	fmt.Print("\033[H")
}
