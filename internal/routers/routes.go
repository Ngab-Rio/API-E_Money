package routers

import (
	"e-money/internal/handlers"
	"e-money/internal/middlewares"
	"e-money/internal/utils"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	authHandler *handlers.AuthHandler,
	accountHandler *handlers.AccountHandler,
	transactionHandler *handlers.TransactionHandler,
	profileHandler *handlers.ProfileHandler,
	debtHandler *handlers.DebtHandler,
	jwtManager utils.JWTManager,
) {
	authRoutes := r.Group("/api/auth")
	AuthRoutes(authRoutes, authHandler)

	accountRoutes := r.Group("/api/account")
	accountRoutes.Use(middlewares.AuthMiddleware(jwtManager))
	AccountRoutes(accountRoutes, accountHandler)

	transactionRoutes := r.Group("/api/transaction")
	transactionRoutes.Use(middlewares.AuthMiddleware(jwtManager))
	TransactionRoutes(transactionRoutes, transactionHandler)

	profileRoutes := r.Group("/api/profile")
	profileRoutes.Use(middlewares.AuthMiddleware(jwtManager))
	ProfileRoutes(profileRoutes, profileHandler)

	debtRoutes := r.Group("/api/debts")
	debtRoutes.Use(middlewares.AuthMiddleware(jwtManager))
	DebtRoutes(debtRoutes, debtHandler)

}
