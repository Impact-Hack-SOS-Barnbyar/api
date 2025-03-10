package config

type AppConfig struct {
	Port uint16
}

func Get() AppConfig {
	return AppConfig{
		Port: 8080,
	}
}
