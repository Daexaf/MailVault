package main

import (
	"log"

	"github.com/daexaf/mailvault/internal/bootstrap"
	"github.com/daexaf/mailvault/internal/handler"
	"github.com/daexaf/mailvault/internal/infrastructure/persistence/gorm"
	branchrepo "github.com/daexaf/mailvault/internal/infrastructure/persistence/gorm"
	"github.com/daexaf/mailvault/internal/route"
	"github.com/daexaf/mailvault/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {

	app, err := bootstrap.New()
	if err != nil {
		log.Fatal(err)
	}

	//repository
	repo := branchrepo.NewBranchRepository(app.DB)
	mailAccountRepo := gorm.NewMailAccountRepository(app.DB)

	//service
	branchService := service.NewBranchService(repo)
	mailAccountService := service.NewMailAccountService(
		mailAccountRepo,
		repo,
	)

	//handler
	branchHandler := handler.NewBranchHandler(branchService)
	mailAccountHandler := handler.NewMailAccountHandler(mailAccountService)

	router := gin.Default()

	api := router.Group("/api/v1")

	branchGroup := api.Group("/branches")
	mailGroup := api.Group("/mail-accounts")

	route.RegisterBranchRoutes(branchGroup, branchHandler)
	route.RegisterMailAccountRoutes(mailGroup, mailAccountHandler)

	if err := router.Run(":" + app.Config.AppPort); err != nil {
		log.Fatal(err)
	}

}
