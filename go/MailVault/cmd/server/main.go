package main

import (
	"log"

	"github.com/daexaf/mailvault/internal/bootstrap"
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

}
