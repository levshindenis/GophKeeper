package config

import (
	"flag"
	"os"
)

type Config struct {
	serverAddress string
	dbAddress     string
}

func (c *Config) GetServerAddress() string {
	return c.serverAddress
}

func (c *Config) GetDBAddress() string {
	return c.dbAddress
}

func (c *Config) SetServerAddress(value string) {
	c.serverAddress = value
}

func (c *Config) SetDBAddress(value string) {
	c.dbAddress = value
}

func (c *Config) ParseFlags() error {
	flag.StringVar(&c.serverAddress, "s", "localhost:8080", "address and port to run server")
	flag.StringVar(&c.dbAddress, "d", "", "db address")
	flag.Parse()

	if envServerAddress, in := os.LookupEnv("SERVER_ADDRESS"); in {
		c.SetServerAddress(envServerAddress)
	}

	if envDBAddress, in := os.LookupEnv("DB_ADDRESS"); in {
		c.SetDBAddress(envDBAddress)
	}

	return nil
}
