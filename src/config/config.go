package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
	SSLMode  bool
}

//GetConfig icin sad sa
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Host:     "localhost",
			Port:     5432,
			Username: "postgres",
			Password: "kubra",
			Name:     "kubDeneme",
			Charset:  "utf8",
			SSLMode:  false,
		},
	}
}
