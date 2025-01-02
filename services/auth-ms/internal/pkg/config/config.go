package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	ServiceID   string `envconfig:"SERVICE_ID"   mapstructure:"SERVICE_ID"`
	ServiceName string `envconfig:"SERVICE_NAME" mapstructure:"SERVICE_NAME"`

	RESTPort string `envconfig:"REST_PORT" mapstructure:"REST_PORT"`
	GRPCPort string `envconfig:"GRPC_PORT" mapstructure:"GRPC_PORT"`

	PrivateKeyPath string `envconfig:"PRIVATE_KEY_PATH" mapstructure:"PRIVATE_KEY_PATH"`
	PublicKeyPath  string `envconfig:"PUBLIC_KEY_PATH"  mapstructure:"PUBLIC_KEY_PATH"`
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
