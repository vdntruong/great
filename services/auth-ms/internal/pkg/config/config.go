package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	RESTAddress string `mapstructure:"rest_address"`
	GRPCAddress string `mapstructure:"grpc_address"`
}

var Cfg *Config

func Load() error {
	cfgFile := configFilePath()
	viper.SetConfigFile(cfgFile)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		return err
	}

	viper.OnConfigChange(func(e fsnotify.Event) {
		_ = viper.ReadInConfig()
		_ = viper.Unmarshal(&Cfg)
	})
	viper.WatchConfig()
	return nil
}

func configFilePath() string {
	if f := os.Getenv(EnvConfigFileKey); len(f) > 0 {
		return f
	}

	return DevConfigFile
}
