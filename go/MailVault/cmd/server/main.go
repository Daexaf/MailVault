package main

import (
	"log"

	"github.com/daexaf/mailvault/internal/bootstrap"
	"github.com/daexaf/mailvault/internal/handler"
	branchrepo "github.com/daexaf/mailvault/internal/infrastructure/persistence/gorm"
	"github.com/daexaf/mailvault/internal/route"
	"github.com/daexaf/mailvault/internal/service"
	"github.com/gin-gonic/gin"
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

	// app, err := bootstrap.New()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// app.Logger.Info("MailVault is starting...")

	// app.Logger.Info(
	// 	"Configuration loaded",
	// 	"Application", app.Config.AppName,
	// 	"Port", app.Config.AppPort,
	// )

	// app.Logger.Info("Database connected successfully")

	// boot, err := bootstrap.New()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// repo := gorm.NewBranchRepository(boot.DB)

	// service := service.NewBranchService(repo)

	// branch, err := service.Create(dto.CreateBranchRequest{
	// 	Code: "BDG",
	// 	Name: "Bandung",
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// }

	app, err := bootstrap.New()
	if err != nil {
		log.Fatal(err)
	}

	//repository
	repo := branchrepo.NewBranchRepository(app.DB)
	//service
	branchService := service.NewBranchService(repo)
	//handler
	branchHandler := handler.NewBranchHandler(branchService)

	router := gin.Default()

	api := router.Group("/api/v1")

	branchGroup := api.Group("/branches")

	route.RegisterBranchRoutes(branchGroup, branchHandler)
	// branch, err := branchService.Create(dto.CreateBranchRequest{
	// 	Code: "BDG",
	// 	Name: "Bandung",
	// })

	if err := router.Run(":" + app.Config.AppPort); err != nil {
		log.Fatal(err)
	}

	// app.Logger.Info(
	// 	"Branch created",
	// 	"id", branch.ID,
	// 	"code", branch.Code,
	// )
}
