package playtika

import (
        "fmt"
        "log"
)

type Config struct {
	username string
	password string
	endpoint string

}

/*
* Builds a client object for this config
 */
func (c *Config) validateAndConnect() error {
	log.Println("Config")
	return nil
}

