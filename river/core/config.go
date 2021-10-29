package core

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type internalConfig struct {
	Dev_mode  bool   `env:"DEV_MODE" envDefault:"false"`
	Log_level string `env:"LOG_LEVEL" envDefault:"debug"`
	Port      uint32 `env:"PORT" envDefault:"1935"`
}

type Config struct {
	Dev_mode  bool
	Log_level log.Level
	Port      uint32
}

func NewConfig() (*Config, error) {
	var err error

	internalCfg := internalConfig{}
	if err = env.Parse(&internalCfg); err != nil {
		return nil, err
	}
	log.Debug(internalCfg)

	cfg := Config{}
	cfg.Dev_mode = internalCfg.Dev_mode
	if cfg.Log_level, err = log.ParseLevel(internalCfg.Log_level); err != nil {
		return nil, err
	}
	cfg.Port = internalCfg.Port

	return &cfg, nil
}
