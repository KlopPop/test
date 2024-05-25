package config

import (
	"log"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env               string
	ServerAddress     string
	ServerTimeout     time.Duration
	ServerIdleTimeout time.Duration
	DBHost            string
	DBPort            int
	DBUser            string
	DBPassword        string
	DBName            string
	Delay             int
}

// временно
func setLocalEnv() {
	os.Setenv("APP_ENV", "local")
	os.Setenv("APP_SERVERADDRESS", "localhost:8080")
	os.Setenv("APP_SERVERTIMEOUT", "5s")
	os.Setenv("APP_SERVERIDLETIMEOUT", "60s")
	os.Setenv("APP_DELAY", "5")
}

func MustLoad() Config {

	setLocalEnv()

	cfg := Config{}

	if err := envconfig.Process("APP", &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return cfg
}
