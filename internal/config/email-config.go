package config

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

var smtp any

func LoadSMTPConfig() *SMTPConfig {
	return &SMTPConfig{
		Host:     "smtp.gmail.com",
		Port:     587,
		Username: "burmaud1@gmail.com",
		Password: "", // use env var in prod
		From:     "burmaud1@gmail.com",
	}
}
