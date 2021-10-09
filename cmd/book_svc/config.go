package main

import (
	"fmt"

	"github.com/namsral/flag"
)

var cfg *AppConfig

type AppConfig struct {
	IsDev    bool
	HttpPort string
	DBConfig
	JaegerConfig
}

func GetConfig() AppConfig {
	if cfg == nil {
		cfg = &AppConfig{}
		flag.BoolVar(&cfg.IsDev, "IS_DEV", false, "Application env")

		flag.StringVar(&cfg.HttpPort, "HTTP_PORT", ":5000", "port for http server")

		// DBConfig
		flag.StringVar(&cfg.DBConfig.Host, "POSTGRES_HOST", "", "")
		flag.StringVar(&cfg.DBConfig.Port, "POSTGRES_PORT", "30100", "")
		flag.StringVar(&cfg.DBConfig.DBName, "POSTGRES_DB", "chat", "")
		flag.StringVar(&cfg.DBConfig.User, "POSTGRES_USER", "root", "")
		flag.StringVar(&cfg.DBConfig.Pass, "POSTGRES_PASS", "password", "")

		// JaegerConfig
		flag.StringVar(&cfg.JaegerConfig.ServiceName, "JAEGER_SERVICE_NAME", "", "")
		flag.StringVar(&cfg.JaegerConfig.AgentHost, "JAEGER_HOST", "", "")
		flag.StringVar(&cfg.JaegerConfig.AgentPort, "JAEGER_PORT", "", "")

		flag.Parse()
	}
	return *cfg
}

// DBConfig config for postgres database
type DBConfig struct {
	Host   string
	Port   string
	DBName string
	User   string
	Pass   string
}

func (c DBConfig) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		c.Host, c.User, c.Pass, c.DBName, c.Port)
}

type JaegerConfig struct {
	ServiceName string
	AgentHost   string
	AgentPort   string
}
