package monitor

import (
	"github.com/zhangpetergo/LiveStreamRecorder/app/processor"
	"github.com/zhangpetergo/LiveStreamRecorder/app/resolver/douyin"
	"github.com/zhangpetergo/LiveStreamRecorder/foundation/logger"
	"github.com/zhangpetergo/LiveStreamRecorder/foundation/urlutil"
)

// 维护已启动录制的直播流
var recordingStreams = make(map[string]bool)

// 维护 url 和对应的直播状态
var liveStatus = make(map[string]bool)

// 检查直播是否在线
func isLiveOnline(url string) bool {

	platformType := urlutil.GetPlatformFromURL(url)
	switch platformType {
	case "douyin":
		live, err := douyin.CheckLiveStream(url)
		if err != nil {
			return false
		}
		if live {
			return true
		}
		return false
	default:
		return false
	}
}

// MonitorStreams 轮询直播状态，每隔 `interval` 秒检测一次
func MonitorStreams(streams []string) {

	for _, stream := range streams {

		if !recordingStreams[stream] {
			live := isLiveOnline(stream)
			if live {
				// 直播开播，启动录制
				go func() {
					defer func() {
						if r := recover(); r != nil {
							logger.Log.Errorw("ProcessStream", "url", stream, "recover", r)
						}
					}()

					liveStatus[stream] = true
					err := processor.ProcessStream(stream)
					if err != nil {
						logger.Log.Errorw("ProcessStream", "url", stream, "err", err)
					}
					// 直播下播，更新录制状态为false
					recordingStreams[stream] = false
				}()
			}
			recordingStreams[stream] = false
		}
	}

}
