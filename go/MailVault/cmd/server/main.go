package main

import (
	"log"

	"github.com/daexaf/mailvault/internal/bootstrap"
	"github.com/daexaf/mailvault/internal/dto"
	"github.com/daexaf/mailvault/internal/infrastructure/persistence/gorm"
	"github.com/daexaf/mailvault/internal/service"
)

func main() {
	// fmt.Println("MailVault is starting...")
	// cfg, err := config.Load()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("==========================")
	// fmt.Println(cfg.AppName)
	// fmt.Println(cfg.AppPort)
	// fmt.Println(cfg.StoragePath)
	// fmt.Println("==========================")

	app, err := bootstrap.New()
	if err != nil {
		log.Fatal(err)
	}

	app.Logger.Info("MailVault is starting...")

	app.Logger.Info(
		"Configuration loaded",
		"Application", app.Config.AppName,
		"Port", app.Config.AppPort,
	)

	app.Logger.Info("Database connected successfully")

	boot, err := bootstrap.New()
	if err != nil {
		log.Fatal(err)
	}

	repo := gorm.NewBranchRepository(boot.DB)

	service := service.NewBranchService(repo)

	err = service.Create(dto.CreateBranchRequest{
		Code: "BDG",
		Name: "Bandung",
	})

	if err != nil {
		log.Fatal(err)
	}
}
