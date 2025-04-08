package config

import (
	"log"
	"os"
	"strings"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
		Insecure bool   `mapstructure:"insecure"`
	}
}

func LoadConfig() (*Config, error) {
	env := os.Getenv("env")
	if env == "" {
		env = "dev"
	}

	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Fatalf("error loading the .env file: %s", err.Error())
	}

	// doesnt map env vars to struct fields unless the keys match exactly
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // database.host -> DATABASE_HOST

	// bind env vars to nested fields
	_ = viper.BindEnv("database.host")
	_ = viper.BindEnv("database.port")
	_ = viper.BindEnv("database.user")
	_ = viper.BindEnv("database.password")
	_ = viper.BindEnv("database.name")
	_ = viper.BindEnv("database.insecure")

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
