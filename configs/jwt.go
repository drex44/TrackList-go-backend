package configs

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Represents database server and credentials
type JWTConfig struct {
	EncryptionKey string
}

// Read and parse the configuration file
func (c *JWTConfig) Read() {
	if _, err := toml.DecodeFile("configs/jwt.toml", &c); err != nil {
		log.Fatal(err)
	}
}
