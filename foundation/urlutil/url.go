package urlutil

import (
	"net/url"
	"strings"
)

// PlatformType 代表直播平台
type PlatformType string

const (
	PlatformUnknown  PlatformType = "unknown"
	PlatformDouyin   PlatformType = "douyin"
	PlatformBilibili PlatformType = "bilibili"
	PlatformDouyu    PlatformType = "douyu"
	PlatformHuya     PlatformType = "huya"
	PlatformTwitch   PlatformType = "twitch"
	PlatformYouTube  PlatformType = "youtube"
)

// GetPlatformFromURL 解析 URL，返回所属平台
func GetPlatformFromURL(streamURL string) PlatformType {
	parsedURL, err := url.Parse(streamURL)
	if err != nil {
		return PlatformUnknown
	}

	host := parsedURL.Host

	switch {
	case strings.Contains(host, "douyin.com"):
		return PlatformDouyin
	case strings.Contains(host, "bilibili.com"):
		return PlatformBilibili
	case strings.Contains(host, "douyu.com"):
		return PlatformDouyu
	case strings.Contains(host, "huya.com"):
		return PlatformHuya
	case strings.Contains(host, "twitch.tv"):
		return PlatformTwitch
	case strings.Contains(host, "youtube.com") || strings.Contains(host, "youtu.be"):
		return PlatformYouTube
	default:
		return PlatformUnknown
	}
}
