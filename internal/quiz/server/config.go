package server

// Config 
type Config struct {
	BindAddress string
	LogLevel	string
}

func NewConfig() *Config{
	return &Config{
		BindAddress: ":8080",
		LogLevel:	"info",
	}
}