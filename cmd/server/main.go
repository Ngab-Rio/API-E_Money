package main

import (
	"e-money/internal/config"
	"e-money/internal/database"
	"e-money/internal/handlers"
	"e-money/internal/repository"
	"e-money/internal/routers"
	"e-money/internal/services"
	"e-money/internal/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	cfg := config.Load()
	database.Connect(cfg.DBDSN())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World",
		})
	})
	jwtManager := utils.NewJWTManager(
		cfg.JWTSecret,
		cfg.JWTIssuer,
		1*time.Hour, // expired token
	)

	// AUTH
	authRepo := repository.NewAuthRepository(database.DB)
	accountRepo := repository.NewAccountRepository(database.DB)
	authService := services.NewAuthService(authRepo, accountRepo, jwtManager)
	authHandler := handlers.NewAuthHandler(authService)

	// ACCOUNTS
	accountService := services.NewAccountService(accountRepo)
	accountHandler := handlers.NewAccountHandler(accountService)

	// TRANSACTIONS
	transactionRepo := repository.NewTransactionRepository(database.DB)
	transactionService := services.NewTransactionService(transactionRepo, accountRepo)
	transactionHandler := handlers.NewTransactionHandler(transactionService, accountService)

	// PROFILES
	profileRepo := repository.NewProfileRepository(database.DB)
	profileService := services.NewProfileService(profileRepo)
	profileHandler := handlers.NewProfileHandler(profileService)

	// DEBTS
	debtRepo := repository.NewDebtRepository(database.DB)
	debtService := services.NewDebtService(debtRepo)
	debtHandler := handlers.NewDebtHandler(debtService)

	routers.SetupRoutes(router, authHandler, accountHandler, transactionHandler, profileHandler, debtHandler, jwtManager)

	router.Run(":7890")
}
