package util

import "os"

// Config struct defines the config structure
type Config struct {
	Mongo       MongoConfig
	Port        string
	CheckURL    string
	CheckString string
}

// MongoConfig has config values for Mongo
type MongoConfig struct {
	Host  string
	DB    string
	Table string
}

// NewConfig parses config file and return Config struct
func NewConfig() *Config {
	config := &Config{
		Mongo: MongoConfig{
			Host:  "127.0.0.1:27017",
			DB:    "proxy_pool",
			Table: "proxies",
		},
		Port:        "8080",
		CheckURL:    "http://www.httpbin.org/get",
		CheckString: "headers",
	}
	if os.Getenv("MONGO_HOST") != "" {
		config.Mongo.Host = os.Getenv("MONGO_HOST")
	}
	if os.Getenv("CHECK_URL") != "" {
		config.CheckURL = os.Getenv("CHECK_URL")
	}
	if os.Getenv("CHECK_STRING") != "" {
		config.CheckString = os.Getenv("CHECK_STRING")
	}

	return config
}
