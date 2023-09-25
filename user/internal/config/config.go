package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Mysql struct {
		Host     string
		Port     int
		Config   string
		Db       string
		User     string
		Password string
	}
	Email struct {
		ValidEmail string
		SmtpHost   string
		SmtpEmail  string
		SmtpPass   string
	}
}
