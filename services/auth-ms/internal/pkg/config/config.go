package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	RESTAddress string `mapstructure:"rest_address"`
	GRPCAddress string `mapstructure:"grpc_address"`

	PrivateKeyPath string `mapstructure:"private_key_path"`
	PublicKeyPath  string `mapstructure:"public_key_path"`

	UserGRPCAddress string `mapstructure:"user_grpc_address"`
}

func Load() (*Config, error) {
	cfgFile := configFilePath()
	viper.SetConfigFile(cfgFile)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config, %v", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("[CONFIG ERROR] failed to read config change, err: %v\n", err)
			return
		}
		if err := viper.Unmarshal(&cfg); err != nil {
			fmt.Printf("[CONFIG ERROR] failed to unmarshal config change, err: %v\n", err)
		}
	})

	return &cfg, nil
}

func configFilePath() string {
	if f := os.Getenv(EnvConfigFileKey); len(f) > 0 {
		return f
	}
	return DevConfigFile
}
