package core

import (
	"fmt"
	"os"
	"strconv"
)

const default_port = uint32(8000)

type Config struct {
	Dev_mode         bool   // default: false
	Disable_register bool   // default: false
	Port             uint32 // default: 8000
	DB_URI           string // mandatory
	Redis_URI        string // mandatory
	Redis_password   string // optional
	AMQP_URI         string // mandatory
	S3_URI           string // mandatory
	S3_access_key    string // mandatory
	S3_secret_key    string // mandatory
	S3_disable_tls   bool   // default: false
}

func ParseConfig() (*Config, error) {
	var config Config
	var err error

	config.Dev_mode = os.Getenv("NIENNA_DEV") == "true"
	config.Disable_register = os.Getenv("DISABLE_NIENNA_REGISTER") == "true"
	if config.Port, err = getPort(); err != nil {
		return &Config{}, err
	}
	if config.DB_URI = os.Getenv("DB_URI"); config.DB_URI == "" {
		return &Config{}, fmt.Errorf("missing mandatory Env param 'DB_URI'")
	}
	if config.Redis_URI = os.Getenv("REDIS_URI"); config.Redis_URI == "" {
		return &Config{}, fmt.Errorf("missing mandatory Env param 'REDIS_URI'")
	}
	config.Redis_password = os.Getenv("REDIS_PASSWORD")
	if config.AMQP_URI = os.Getenv("AMQP_URI"); config.AMQP_URI == "" {
		return &Config{}, fmt.Errorf("missing mandatory Env param 'AMQP_URI'")
	}
	if config.S3_URI = os.Getenv("S3_URI"); config.S3_URI == "" {
		return &Config{}, fmt.Errorf("missing mandatory Env param 'S3_URI'")
	}
	if config.S3_access_key = os.Getenv("S3_ACCESS_KEY"); config.S3_access_key == "" {
		return &Config{}, fmt.Errorf("missing mandatory Env param 'S3_ACCESS_KEY'")
	}
	if config.S3_secret_key = os.Getenv("S3_SECRET_KEY"); config.S3_secret_key == "" {
		return &Config{}, fmt.Errorf("missing mandatory Env param 'S3_SECRET_KEY'")
	}
	config.S3_disable_tls = os.Getenv("S3_DISABLE_TLS") == "true"

	return &config, nil
}

func getPort() (uint32, error) {
	var err error
	if os.Getenv("PORT") == "" {
		return default_port, nil
	}
	parsedPort, err := strconv.ParseInt(os.Getenv("PORT"), 10, 64)
	if err != nil {
		return default_port, err
	}
	return uint32(parsedPort), nil
}
