package singleton

import "sync"

type Config struct {
	APIKey string
}

var (
	instance *Config
	once     sync.Once
)

func GetInstance() *Config {
	once.Do(func() {
		instance = &Config{
			APIKey: "your-api-key",
		}
	})

	return instance
}
