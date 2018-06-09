package configs

import (
	"log"
	"github.com/BurntSushi/toml"
)

// Represents database server and credentials
type MongoConfig struct {
	Server   string
	Database string
}

// Read and parse the configuration file
func (c *MongoConfig) Read() {
	if _, err := toml.DecodeFile("mongo.toml", &c); err != nil {
		log.Fatal(err)
	}
}