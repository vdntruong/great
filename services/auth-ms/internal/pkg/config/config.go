package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	ServiceID   string `mapstructure:"SERVICE_ID" envconfig:"SERVICE_ID"`
	ServiceName string `mapstructure:"SERVICE_NAME" envconfig:"SERVICE_NAME"`

	RESTPort string `mapstructure:"REST_PORT" envconfig:"REST_PORT"`
	GRPCPort string `mapstructure:"GRPC_PORT" envconfig:"GRPC_PORT"`

	PrivateKeyPath string `mapstructure:"PRIVATE_KEY_PATH" envconfig:"PRIVATE_KEY_PATH"`
	PublicKeyPath  string `mapstructure:"PUBLIC_KEY_PATH" envconfig:"PUBLIC_KEY_PATH"`
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
	return DefaultConfigFile
}
