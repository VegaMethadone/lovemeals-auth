package config

type Config struct {
	Version int    `json:"version"`
	Env     string `json:"env"`
	Network struct {
		Address      string `json:"address"`
		Port         string `json:"port"`
		WriteTimeout int    `json:"writeTimeout"`
		ReadTimeout  int    `json:"readTimeout"`
	} `json:"network"`
	Security struct {
		SSLCertificate         string `json:"ssl_certificate"`
		SSLKey                 string `json:"ssl_key"`
		AuthenticationRequired bool   `json:"authentication_required"`
	} `json:"security"`
	Postgres struct {
		Host         string `json:"host"`
		Port         int    `json:"port"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		DatabaseName string `json:"database_name"`
		SSLMode      string `json:"sslmode"`
	} `json:"postgres"`
}
