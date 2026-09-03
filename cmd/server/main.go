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
	emailRepo := gorm.NewEmailRepository(app.DB)
	attachmentRepo := gorm.NewAttachmentRepository(app.DB)
	emailService := service.NewEmailService(
		emailRepo,
	)

	emailHandler := handler.NewEmailHandler(
		emailService,
	)

	//service
	branchService := service.NewBranchService(repo)
	mailAccountService := service.NewMailAccountService(
		mailAccountRepo,
		repo,
		emailRepo,
		attachmentRepo,
	)

	//handler
	branchHandler := handler.NewBranchHandler(branchService)
	mailAccountHandler := handler.NewMailAccountHandler(mailAccountService)

	router := gin.Default()

	api := router.Group("/api/v1")
	emailGroup := api.Group("/emails")

	route.RegisterEmailRoutes(
		emailGroup,
		emailHandler,
	)

	branchGroup := api.Group("/branches")
	mailGroup := api.Group("/mail-accounts")

	route.RegisterBranchRoutes(branchGroup, branchHandler)
	route.RegisterMailAccountRoutes(mailGroup, mailAccountHandler)

	// err = imap.TestConnection(
	// 	"imap.gmail.com",
	// 	993,
	// 	"fitsan64@gmail.com",
	// 	"ixrl lzuh qtti taoe",
	// )

	if err != nil {
		log.Fatal(err)
	}

	if err := router.Run(":" + app.Config.AppPort); err != nil {
		log.Fatal(err)
	}

}
