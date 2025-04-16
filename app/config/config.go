// Package config 管理配置文件相关的操作
package config

import (
	"github.com/pkg/errors"
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
	SavePath               string `mapstructure:"save_path"`
	PollIntervalSeconds    int    `mapstructure:"poll_interval_seconds"`
	EnableSegmenting       bool   `mapstructure:"enable_segmenting"`
	SegmentDurationSeconds int    `mapstructure:"segment_duration_seconds"`
}

func SetConfigPath(path string) {
	configPath = path
}

func setDefaults() {
	viper.SetDefault("save_path", "./downloads")
	viper.SetDefault("poll_interval_seconds", 60)
	viper.SetDefault("enable_segmenting", true)
	viper.SetDefault("segment_duration_seconds", 1800)
}

func GetConfig() (*Config, error) {

	once.Do(func() {

		if configPath == "" {
			loadErr = errors.Wrap(loadErr, "config path not set")
			return
		}

		// -------------------------------------------------------------------------
		setDefaults()

		// -------------------------------------------------------------------------
		viper.SetConfigFile(configPath)
		if err := viper.ReadInConfig(); err != nil {
			loadErr = errors.Wrap(loadErr, "viper.ReadInConfig")
			return
		}

		// -------------------------------------------------------------------------
		if err := viper.Unmarshal(&config); err != nil {
			loadErr = errors.Wrap(loadErr, "viper.Unmarshal")
			return
		}
	})

	return config, loadErr
}
