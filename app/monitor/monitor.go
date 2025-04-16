package monitor

import (
	"github.com/zhangpetergo/LiveStreamRecorder/app/processor"
	"github.com/zhangpetergo/LiveStreamRecorder/app/resolver/douyin"
	"github.com/zhangpetergo/LiveStreamRecorder/foundation/logger"
	"github.com/zhangpetergo/LiveStreamRecorder/foundation/urlutil"
)

// 维护已启动录制的直播流
var recordingStreams = make(map[string]bool)

// 检查直播是否在线
func isLiveOnline(url string) bool {

	platformType := urlutil.GetPlatformFromURL(url)
	switch platformType {
	case urlutil.PlatformDouyin:
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

// Listen 轮询直播状态，每隔 `interval` 秒检测一次
func Listen(streams []string) {

	for _, stream := range streams {

		if !recordingStreams[stream] {
			live := isLiveOnline(stream)
			if live {
				logger.Log.Infow("检测到直播开播", "url", stream)
				// 直播开播，启动录制
				go func() {
					// -------------------------------------------------------------------------
					defer func() {
						if r := recover(); r != nil {
							logger.Log.Errorw("ProcessStream", "url", stream, "recover", r)
							// 如果发生错误，更新直播状态为 false
							recordingStreams[stream] = false
						}
					}()

					// -------------------------------------------------------------------------
					err := processor.ProcessStream(stream)
					if err != nil {
						logger.Log.Errorf("ProcessStream url: %s %+v ", stream, err)
						// 如果发生错误，更新直播状态为 false
						recordingStreams[stream] = false
					}

					// -------------------------------------------------------------------------
					// 直播下播，更新录制状态为false
					recordingStreams[stream] = false
				}()
			}
			recordingStreams[stream] = false
		}
	}

}
