package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env      string         `yaml:"env" env-default:"local"`
	HTTP     HTTPConfig     `yaml:"http"`
	Postgres PostgresConfig `yaml:"postgres"`
}

type HTTPConfig struct {
	Port string `yaml:"port"`
}

type PostgresConfig struct {
	Url string `yaml:"url"`
}

func MustLoad() *Config {

	path := fetchConfigPath()

	if path == "" {
		panic("config path is empty")
	}

	return MustLoadPath(path)
}

func MustLoadPath(path string) *Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file doesn't exist" + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "cfg", "", "path to config file")
	flag.Parse()

	if res == "" {
		if err := godotenv.Load(".env"); err != nil {
			panic("failed to load .env file: " + err.Error())
		}
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
