package setup

import "github.com/kelseyhightower/envconfig"

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
	WebPort string `default:"8080"`
}

func newConfig() config {
	var c config
	if err := envconfig.Process("WHT", &c); err != nil {
		return c
	}
	return c
}

func (c *config) IsLocal() bool {
	return c.Env == "local"
}
