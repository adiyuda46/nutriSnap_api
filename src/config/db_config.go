// config/config.go
package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Dns string
}

func CreateConnectionDB() *Config {
	config := new(Config)

	var (
		server = viper.GetString("nutrisnap.server")
		port   = viper.GetString("nutrisnap.portdb")
		user   = viper.GetString("nutrisnap.user")
		pass   = viper.GetString("nutrisnap.password")
		scheme = viper.GetString("nutrisnap.scheme")
	)
	// Validasi data konfigurasi
	if server == "" || port == "" || user == "" || pass == "" || scheme == "" {
		fmt.Println("⚠️  Konfigurasi tidak lengkap! Pastikan semua field di config.json terisi.")
		return nil
	}

	postgresInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		server,
		port,
		user,
		pass,
		scheme)

	config.Dns = postgresInfo
	return config
}
