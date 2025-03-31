// monitor_test.go
package monitor

import (
	"github.com/stretchr/testify/assert"
	"github.com/zhangpetergo/LiveStreamRecorder/foundation/logger"
	"testing"
	"time"
)

// 对于 url 正确且在线的直播，应该返回 true
func TestIsLiveOnlineReturnsTrueForValidStream(t *testing.T) {
	logger.InitLogger()

	//------------------------------------------------
	// douyin url
	url := "https://live.douyin.com/398706122526"
	result := isLiveOnline(url)

	assert.True(t, result)
}

// 如果 url 正确但不在线，应该返回 false
func TestIsLiveOnlineReturnsFalseForOfflineStream(t *testing.T) {
	logger.InitLogger()
	//------------------------------------------------
	// douyin url
	url := "https://live.douyin.com/944469102465"
	result := isLiveOnline(url)

	assert.False(t, result)
}

// 如果 url 格式不正确，应该返回 false
func TestIsLiveOnlineReturnsFalseForInvalidURL(t *testing.T) {
	logger.InitLogger()
	//------------------------------------------------
	url := "invalid-url"
	result := isLiveOnline(url)

	assert.False(t, result)
}

// stream 列表为空时，Listen 函数应该不执行任何操作
func TestListenHandlesEmptyStreamList(t *testing.T) {
	logger.InitLogger()

	var streams []string
	Listen(streams)

	assert.Equal(t, 0, len(streams))
}

// 测试 Listen 函数是否正确处理多个流
// 这里目前测试正常，后续应该考虑更优雅的方法测试
func TestListenHandlesMultipleStreams(t *testing.T) {
	logger.InitLogger()

	//------------------------------------------------
	// douyin url
	streams := []string{
		"https://live.douyin.com/398706122526",
		"https://live.douyin.com/695496496290",
	}
	Listen(streams)

	time.Sleep(15 * time.Second)
}
