package router

import (
	"golang-e-wallet-rest-api/internal/router/middlewares"
	"golang-e-wallet-rest-api/internal/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *server.HandlerOps) http.Handler {
	g := gin.New()
	g.ContextWithFallback = true

	g.Use(gin.Recovery(), middlewares.LoggerMiddleware, middlewares.ErrorMiddleware)
	g.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "page not found"})
	})

	g.POST("/register", h.UserController.Register)
	g.POST("/login", h.UserController.Login)
	g.POST("/forgot-password", h.UserController.ForgotPassword)
	g.POST("/reset-password", h.UserController.ResetPassword)

	user := g.Group("/user", middlewares.JWTAuth)
	user.POST("/transfer", h.TransactionController.Transfer)
	user.POST("/top-up", h.TransactionController.TopUp)
	user.GET("/transactions", h.TransactionController.GetAllTransactions)
	user.GET("/details", h.UserController.GetDetails)

	return g
}
