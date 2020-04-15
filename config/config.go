package config

import (
	"os"
	"strconv"
)

type DBConfig struct {
	Name string
	Host string
	Port int
	Username string
	Password string
}

type RedisConfig struct {
	Host string
	Port int
	Pass string
	DB int
}

type Config struct {
	Debug bool
	Port int
	Host string
	DB DBConfig
	Redis RedisConfig
}

var currentConfig Config

func GetConfig() Config {
	return	currentConfig
}

func InitConfig() Config {
	key := os.Getenv("ENV")
	if key == "" {
		key = "production"
	}
	mongo := DBConfig{
		Name:     "novelweb",
		Host:     "127.0.0.1",
		Port:     27017,
		Username: "root",
		Password: "",
	}
	redis := RedisConfig{
		Host: "127.0.0.1",
		Port: 6379,
		Pass: "",
		DB: 0,
	}
	switch key {
	case	"production":
		currentConfig = Config{
			Debug: false,
			Port:  8900,
			Host:  "",
			DB:    DBConfig{},
			Redis: RedisConfig{},
		}
		break
	case "dev":
		currentConfig = Config{
			Debug: true,
			Port:  8900,
			Host:  "127.0.0.1",
			DB:    mongo,
			Redis: redis,
		}
		break
	default:
		currentConfig = Config{}
	}
	if key = os.Getenv("DEBUG"); key != "" {
		if b, err := strconv.ParseBool(key); err == nil {
			currentConfig.Debug = b
		}
	}
	if key = os.Getenv("HOST"); key != "" {
		currentConfig.Host = key
	}
	if key = os.Getenv("PORT"); key != "" {
		if i, err := strconv.Atoi(key); err == nil {
			currentConfig.Port = i
		}
	}
	// mongo
	if key = os.Getenv("DB_HOST"); key != "" {
		currentConfig.DB.Host = key
	}
	if key = os.Getenv("DB_PORT"); key != "" {
		if i, err := strconv.Atoi(key); err == nil {
			currentConfig.DB.Port = i
		}
	}
	if key = os.Getenv("DB_USER"); key != "" {
		currentConfig.DB.Username = key
	}
	if key = os.Getenv("DB_PWD"); key != "" {
		currentConfig.DB.Password = key
	}
	if key = os.Getenv("DB_NAME"); key != "" {
		currentConfig.DB.Name = key
	}
	// redis
	if key = os.Getenv("REDIS_HOST"); key != "" {
		currentConfig.Redis.Host = key
	}
	if key = os.Getenv("REDIS_PORT"); key != "" {
		if i, err := strconv.Atoi(key); err == nil {
			currentConfig.Redis.Port = i
		}
	}
	if key = os.Getenv("REDIS_PWD"); key != "" {
		currentConfig.Redis.Pass = key
	}
	if key = os.Getenv("REDIS_DB"); key != "" {
		if i, err := strconv.Atoi(key); err == nil {
			currentConfig.Redis.DB = i
		}
	}
	return currentConfig
}