package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	serverAddress string
	dbAddress     string
	serverKey     string
}

func (c *Config) GetServerAddress() string {
	return c.serverAddress
}

func (c *Config) GetDBAddress() string {
	return c.dbAddress
}

func (c *Config) GetServerKey() string {
	return c.serverKey
}

func (c *Config) SetServerAddress(value string) {
	c.serverAddress = value
}

func (c *Config) SetDBAddress(value string) {
	c.dbAddress = value
}

func (c *Config) SetServerKey(value string) {
	c.serverKey = value
}

func (c *Config) Parse() error {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	if envServerAddress, in := os.LookupEnv("SERVER_ADDRESS"); in {
		c.SetServerAddress(envServerAddress)
	}

	if envDBAddress, in := os.LookupEnv("DB_ADDRESS"); in {
		c.SetDBAddress(envDBAddress)
	}

	if envServerKey, in := os.LookupEnv("SERVER_KEY"); in {
		c.SetServerKey(envServerKey)
	}

	return nil
}
