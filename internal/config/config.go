package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env     string         `yaml:"env" env-default:"local"`
	Storage PostgresConfig `yaml:"postgres" env-requred:"true"`
	GRPC    GRPCConfig     `yaml:"grpc" env-requred:"true"`
}

type PostgresConfig struct {
	Host     string `yaml:"host" env-requred:"true"`
	Port     int    `yaml:"port" env-requred:"true"`
	User     string `yaml:"user" env-requred:"true"`
	Password string `yaml:"password" env-requred:"true"`
	DBName   string `yaml:"dbname" env-requred:"true"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var configPath string

	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	if configPath == "" {
		configPath = os.Getenv("CONFIG_PATH")
	}

	return configPath
}
