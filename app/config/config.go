package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var (
	once       sync.Once
	config     *Config
	configPath string
	loadErr    error
)

type Config struct {
	SavePath               string `yaml:"save_path"`
	PollIntervalSeconds    int    `yaml:"poll_interval_seconds"`
	EnableSegmenting       bool   `yaml:"enable_segmenting"`
	SegmentDurationSeconds int    `yaml:"segment_duration_seconds"`
}

func SetConfigPath(path string) {
	configPath = path
}

func GetConfig() (*Config, error) {
	once.Do(func() {
		// 使用 viper 读取配置文件
		viper.SetConfigFile(configPath)
		if err := viper.ReadInConfig(); err != nil {
			loadErr = fmt.Errorf("viper.ReadInConfig: %w", err)
			return
		}
		var config Config
		if err := viper.Unmarshal(&config); err != nil {
			loadErr = fmt.Errorf("viper.Unmarshal: %w", err)
			return
		}
	})

	return config, loadErr
}
