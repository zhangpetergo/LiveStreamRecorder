// Package recorder  是用来录制直播流的
package recorder

import (
	"bufio"
	"fmt"
	"os/exec"
	"time"
)

// Record 根据 输入的 url 使用 ffmpeg 下载直播流
func Record(data map[string]interface{}) error {

	// 直播流 url
	url, ok := data["url"].(string)
	if !ok {
		return fmt.Errorf("url 不存在")
	}

	// 作者名
	name, ok := data["name"].(string)
	if !ok {
		return fmt.Errorf("name 不存在")
	}

	// 生成的文件名 name_日期_时间.ts
	outputFile := fmt.Sprintf("%s_%s.ts", name, time.Now().Format("2006-01-02_15-04-05"))

	userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/115.0"

	// 构造 ffmpeg 命令

	cmd := exec.Command(
		"./ffmpeg",
		"-y",            // 覆盖输出文件
		"-v", "verbose", // 输出详细信息
		"-hide_banner", // 隐藏 ffmpeg 的 banner
		"-user_agent", "\""+userAgent+"\"",
		"-protocol_whitelist", "rtmp,crypto,file,http,https,tcp,tls,udp,rtp,httpproxy",
		"-thread_queue_size", "1024",
		"-analyzeduration", "40000000",
		"-probesize", "20000000",
		"-fflags", "+discardcorrupt",
		"-re", "-i",
		url,
		"-bufsize", "1M", // 设置输入缓冲区大小
		"-sn",                     // 禁用字幕
		"-dn",                     // 禁用数据流
		"-rw_timeout", "50000000", // 设置读写操作的超时时间，如果超过这个时间没有数据传输，则会断开连接，防止ffmpeg无限等待
		"-reconnect_delay_max", "60", // 重连的最大延迟时间
		"-reconnect_streamed",            // 重新连接到流媒体服务器
		"-reconnect_at_eof",              // 在文件结束时重新连接
		"-max_muxing_queue_size", "1024", // 设置最大复用队列大小，解决一些错误有用
		"-correct_ts_overflow", "1", // 修正时间戳溢出
		"-avoid_negative_ts", "1",
		"-loglevel", "error",
		"-map", "0",
		"-c:v", "copy",
		"-c:a", "copy",
		"-t", "10",
		outputFile,
	)

	// 输出下构造的命令
	fmt.Println("输出命令")
	fmt.Println(cmd.String())

	// 获取标准错误管道
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Printf("无法获取 stderr: %v\n", err)
		return err
	}

	// 启动命令
	if err := cmd.Start(); err != nil {
		fmt.Printf("命令启动失败: %v\n", err)
		return err
	}

	// 异步读取标准错误
	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			fmt.Printf("错误: %s\n", scanner.Text())
		}
	}()

	// 等待命令完成
	if err := cmd.Wait(); err != nil {
		fmt.Printf("命令执行失败: %v\n", err)
		return err
	}

	fmt.Println("ffmpeg 处理完成，文件保存为", outputFile)

	return nil
}
