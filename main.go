package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"logging-best-practices/internal/handlers"
	"logging-best-practices/pkg/logger"
	"logging-best-practices/pkg/middleware"
)

func main() {
	// 1. Obtener entorno
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	// 2. Inicializar logger
	logger.Init(env)
	defer logger.Sync()

	log := logger.Get()
	log.Info("starting server", zap.String("env", env))

	// 3. Configurar Gin
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New() // Usamos New() para control total de middlewares

	// 4. Aplicar middlewares (RECORDAD aqui el orden importante)
	r.Use(middleware.Recovery())
	r.Use(middleware.Logger())

	// 5. Rutas
	r.GET("/users/:id", handlers.GetUserHandler)
	r.POST("/orders", handlers.CreateOrderHandler)
	r.GET("/products", handlers.GetProductsHandler)

	// 6. Arrancar servidor en goroutine
	go func() {
		log.Info("server listening", zap.String("port", "8080"))
		if err := r.Run(":8080"); err != nil {
			log.Fatal("server failed", zap.Error(err))
		}
	}()

	// 7. Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("shutting down server...")
}
