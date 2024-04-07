package config

import (
	"flag"
	"log"

	xconfig "github.com/cetnfurkan/core/config"
)

type Config struct {
	Echo        xconfig.Server
	PGSQL       xconfig.Database
	RabbitMQ    xconfig.MQ
	Redis       xconfig.Database
	UserService xconfig.Service
}

var (
	configFile string

	cfg = Config{}
)

func init() {
	flag.StringVar(&configFile, "config", "config.yml", "Config file path")
	flag.Parse()

	err := xconfig.Read(configFile, &cfg)
	if err != nil {
		log.Fatal("Unable to read config", err)
	}
}

// Get returns the config instance for the application.
func Get() Config {
	return cfg
}
