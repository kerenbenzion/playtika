package playtika

type Config struct {
	username string
	password string
	endpoint string

	session *api.Client
}

/*
* Builds a client object for this config
 */
func (c *Config) validateAndConnect() error {
	return nil
}

