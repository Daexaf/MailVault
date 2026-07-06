package database

import (
	"fmt"

	"github.com/daexaf/mailvault/internal/config"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// Connect membuat koneksi ke SQL Server menggunakan GORM.
func Connect(cfg *config.Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf(
		"server=%s;user id=%s;password=%s;database=%s;encrypt=disable;TrustServerCertificate=true",
		cfg.DBServer,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	// fmt.Println("========== CONFIG ==========")
	// fmt.Println("Server   :", cfg.DBServer)
	// fmt.Println("Database :", cfg.DBName)
	// fmt.Println("User     :", cfg.DBUser)
	// fmt.Println("Password :", cfg.DBPassword)
	// fmt.Println("DSN      :", dsn)
	// fmt.Println("============================")

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("connect sql server: %w", err)
	}

	return db, nil
}
