package main

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	// 起動環境切り分け用
	Env string `default:"local"`

	// DB接続設定用
	DBHost string `split_words:"true" default:"localhost"`
	DBPort string `split_words:"true" default:"11111"`
	DBName string `split_words:"true" default:"whtdb"`
	DBUser string `split_words:"true" default:"postgres"`
	DBPass string `split_words:"true" default:"yuckyjuice"`

	// Webサーバ設定用
	WebPort string `split_words:"true" default:"8080"`

	// ストレージ設定用
	ImageContentsBucket string `split_words:"true" default:"image-content-bucket"`

	// トレース設定用
	Trace bool `split_words:"true" default:"false"`
}

func newConfig() config {
	var c config
	if err := envconfig.Process("WHT", &c); err != nil {
		return c
	}
	log.Printf("config:%#+v", c)
	return c
}

func (c *config) String() string {
	return fmt.Sprintf("[CONFIG] (WHT_ENV=%s), (WHT_DB_HOST=%s), (WHT_DB_PORT=%s), (WHT_DB_USER=%s), (WHT_DB_PASS=%s), (WHT_WEB_PORT=%s), (WHT_IMAGE_CONTENTS_BUCKET=%s), (WHT_TRACE=%v)",
		c.Env, c.DBHost, c.DBPort, c.DBUser, c.DBPass, c.WebPort, c.ImageContentsBucket, c.Trace)
}

func (c *config) IsLocal() bool {
	return c.Env == "local"
}
