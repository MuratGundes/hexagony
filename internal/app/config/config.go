package config

import "github.com/kelseyhightower/envconfig"

// Config is a configuration struct that contains
// all environment variables of the app.
type Config struct {
	EnvMode      string `envconfig:"ENVMODE" required:"true" default:"development"`
	ServerPort   string `envconfig:"SERVERPORT" required:"true" default:"8000"`
	DBHost       string `envconfig:"DBHOST" default:"localhost" required:"true"`
	DBPort       string `envconfig:"DBPORT" default:"3306" required:"true"`
	DBUser       string `envconfig:"DBUSER" default:"root" required:"true"`
	DBName       string `envconfig:"DBNAME" default:"hexagony" required:"true"`
	DBPass       string `envconfig:"DBPASS" default:"secret" required:"true"`
	RedisAddress string `envconfig:"REDISADDR"`
	RedisPass    string `envconfig:"REDISPASS"`
}

// Load loads the app the configuration based
// in the environment variables.
func Load() (*Config, error) {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
