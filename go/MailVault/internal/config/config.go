package config

// Config menyimpan seluruh konfigurasi aplikasi.
type Config struct {
	AppName string
	AppPort string

	DBServer   string
	DBName     string
	DBUser     string
	DBPassword string

	StoragePath string
}

// func New() *Config {
// 	return &Config{}
// }
