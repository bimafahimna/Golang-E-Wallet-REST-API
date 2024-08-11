package app

import (
	"context"
	"golang-e-wallet-rest-api/internal/databases"
	"golang-e-wallet-rest-api/internal/router"
	"golang-e-wallet-rest-api/internal/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start() {
	db, err := databases.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	if err != nil {
		log.Fatalf("unable to connect to the database: %v", err)
	}
	log.Println("Database connected successfully")

	initHandler := server.SetupHandler(db)
	route := router.SetupRouter(initHandler)

	const Addr = ":8080"

	srv := http.Server{
		Addr:    Addr,
		Handler: route,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err == nil {
			log.Fatal("Server Shutdown")
		}
	}()

	<-ctx.Done()
	log.Println("timeout of 5 seconds.")

	log.Println("Server exiting")
}
