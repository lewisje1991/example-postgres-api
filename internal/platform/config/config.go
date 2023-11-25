package config

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/spf13/viper"
)

type Config struct {
	Mode       string `mapstructure:"MODE"`
	HostPort   int    `mapstructure:"HOST_PORT"`
	DBURL     string `mapstructure:"DB_URL"`

	viper *viper.Viper
}

// NewConfig creates a new config
func NewConfig() *Config {
	return &Config{
		viper: viper.New(),
	}
}

// LoadConfig loads the config from given path or from system environment variables
func (c *Config) Load(filePath string) error {
	c.viper.SetConfigFile(filePath)

	if err := c.viper.ReadInConfig(); err != nil {
		log.Println("did not load env from file falling back to system env", err)
	}

	c.initDefaultvalues()
	c.viper.AutomaticEnv()

	if err := c.viper.Unmarshal(c); err != nil {
		err = fmt.Errorf("could not unmarshal config: %w", err)
		return err
	}

	return nil
}

func (c *Config) initDefaultvalues() {
	c.viper.SetDefault("MODE", "non-prod")
	c.viper.SetDefault("HOST_PORT", 8080)
}

func (c *Config) Validate() error {
	v := reflect.ValueOf(*c)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() == reflect.String {
			if v.Field(i).String() == "" {
				return errors.New(typeOfS.Field(i).Name + " is empty")
			}
		}
		if v.Field(i).Kind() == reflect.Int {
			if v.Field(i).Int() == 0 {
				return errors.New(typeOfS.Field(i).Name + " is 0")
			}
		}
	}

	return nil
}
