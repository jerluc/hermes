package hermes

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config map[interface{}]interface{}

func LoadConfig(path string) (*Config, error) {
	var config Config
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(contents, &config)
	return &config, err
}

func (c Config) Get(key string, defaultValue ...string) string {
	value, keyExists := c[key]
	if !keyExists {
		if len(defaultValue) == 0 {
			return ""
		} else {
			return defaultValue[0]
		}
	}
	return value.(string)
}

func (c Config) GetConfig(key string) Config {
	value, keyExists := c[key]
	if !keyExists {
		return Config{}
	}
	return value.(Config)
}
