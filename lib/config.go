package hermes

type Config map[string]interface{}

func (c Config) Get(key string, defaultValue...string) string {
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