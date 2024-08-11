package server

import (
	"database/sql"
	"golang-e-wallet-rest-api/internal/controllers"
	"golang-e-wallet-rest-api/internal/pkgs/utils/logger"
	"golang-e-wallet-rest-api/internal/repositories"
	"golang-e-wallet-rest-api/internal/services"
)

type HandlerOps struct {
	UserController        controllers.UserController
	TransactionController controllers.TransactionController
}

func SetupHandler(db *sql.DB) *HandlerOps {
	logrus := logger.NewLogger()
	logger.SetLogger(logrus)

	transactor := repositories.InitTransactor(db)

	userRepository := repositories.NewUserRepository(db)
	transactionRepository := repositories.NewTransactionRepository(db)
	walletRepository := repositories.NewWalletRepository(db)

	userService := services.NewUserService(userRepository, walletRepository, transactor)
	transactionService := services.NewTransactionService(transactionRepository, transactor)

	userController := controllers.NewUserController(userService)
	transactionController := controllers.NewTransactionController(transactionService)

	return &HandlerOps{
		UserController:        userController,
		TransactionController: transactionController,
	}
}
