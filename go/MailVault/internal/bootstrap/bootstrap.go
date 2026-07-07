package bootstrap

import (
	"fmt"
	"log/slog"

	"github.com/daexaf/mailvault/internal/config"
	"github.com/daexaf/mailvault/internal/database"
	"github.com/daexaf/mailvault/internal/logger"

	"gorm.io/gorm"
)

type Bootstrap struct {
	Config *config.Config
	Logger *slog.Logger
	DB     *gorm.DB
}

func New() (*Bootstrap, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	log := logger.New()

	db, err := database.Connect(cfg)

	if err != nil {
		return nil, fmt.Errorf("bootstrap database: %w", err)
	}

	if err := database.Migrate(db); err != nil {
		return nil, fmt.Errorf("database migration: %w", err)
	}

	return &Bootstrap{
		Config: cfg,
		Logger: log,
		DB:     db,
	}, nil
}
