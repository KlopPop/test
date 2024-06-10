package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type OldConfig struct {
	Env               string
	LogLevel          int
	LogFormat         string
	ServerAddress     string
	ServerTimeout     time.Duration
	ServerIdleTimeout time.Duration
	DBHost            string
	DBPort            int
	DBUser            string
	DBPassword        string
	DBName            string
}

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	LogLevel   int    `yaml:"loglevel" env-default:"0"`
	LogFormat  string `yaml:"logformat" env-default:"text"`
	Storage    `yaml:"storage" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}

type Storage struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

/*
	log levels
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 4
	LevelError Level = 8

	log Format
	json = "json"
	text = "text"

*/

func MustLoad() *Config {
	/*
		Setenv()

		// временно
		func Setenv() {
			os.Setenv("APP_ENV", "local")
			os.Setenv("APP_LOGLEVEL", "-4")
			os.Setenv("APP_LOGFORMAT", "json")
			os.Setenv("APP_SERVERADDRESS", "localhost:8081")
			os.Setenv("APP_SERVERTIMEOUT", "4s")
			os.Setenv("APP_SERVERIDLETIMEOUT", "60s")
			os.Setenv("APP_DBHOST", "localhost")
			os.Setenv("APP_DBPORT", "5432")
			os.Setenv("APP_DBUSER", "postgres")
			os.Setenv("APP_DBPASSWORD", "admin")
			os.Setenv("APP_DBNAME", "postgres")
		}

		old_cfg := OldConfig{}
		if err := envconfig.Process("APP", &old_cfg); err != nil {
			log.Fatalf("cannot read config: %s", err)
		}

		os.Setenv("STORAGE_CONFIG",
			fmt.Sprintf("host=%s port=%d user=%s "+
				"password=%s dbname=%s sslmode=disable",
				old_cfg.DBHost, old_cfg.DBPort, old_cfg.DBUser, old_cfg.DBPassword, old_cfg.DBName))
	*/

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		configPath = "C:/UserData/Projects/marketplace/product/config/config.yml"
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	os.Setenv("STORAGE_CONFIG",
		fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			cfg.Storage.Host,
			cfg.Storage.Port,
			cfg.Storage.User,
			cfg.Storage.Password,
			cfg.Storage.DBName))

	return &cfg
}
