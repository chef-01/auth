package main

import (
	"auth/common/config"
	"auth/common/database"
	"auth/modules/user/di"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Load application config
	cfg := config.LoadConfig()

	// Initialize MongoDB connection
	dbClient := database.InitMongoDB()
	defer func() {
		if err := dbClient.Disconnect(context.Background()); err != nil {
			log.Fatalf("Error disconnecting MongoDB: %v", err)
		}
	}()

	// Select the database
	db := dbClient.Database(cfg.DatabaseName)

	// Initialize DI container and get router
	router := di.Execute(db, cfg)

	// Configure HTTP server
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		fmt.Printf("🚀 Server is running at http://localhost:%s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	<-done
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed: %+v", err)
	}

	log.Println("Server exited properly 🚀")
}
