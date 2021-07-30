package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Params struct {
	Build         bool   `json:"build_images"`
	Disable_tls   bool   `json:"disable_tls"`
	Dev_mode      bool   `json:"dev_mode"`
	Enable_TTY    bool   `json:"enable_tty"`
	DB_user       string `json:"db_user"`
	DB_password   string
	DB_database   string `json:"db_database"`
	Redis_uri     string `json:"redis_uri"`
	S3_uri        string `json:"s3_uri"`
	S3_key        string `json:"s3_key"`
	S3_secret     string `json:"s3_secret"`
	AMQP_user     string `json:"amqp_user"`
	AMQP_password string
	LB_enable_tls bool `json:"lb_tls"`
}

func ParseParams(path string) (*Params, error) {
	var env Params
	envfile, err := os.Open("env.json")
	if err != nil {
		return nil, err
	}
	byteValue, err := ioutil.ReadAll(envfile)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(byteValue, &env); err != nil {
		return nil, err
	}
	env.DB_password = "TODO"
	env.AMQP_password = "TODO"
	return &env, nil
}
