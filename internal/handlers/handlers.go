package handlers

import (
	"errors"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"logging-best-practices/pkg/logger"
)

func GetUserHandler(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.Param("id")

	logger.Info(ctx, "fetching user from database",
		zap.String("user_id", userID),
	)

	time.Sleep(time.Duration(30+rand.Intn(50)) * time.Millisecond)

	if rand.Float32() < 0.3 {
		err := errors.New("connection timeout")
		logger.Error(ctx, "database error",
			zap.String("user_id", userID),
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		return
	}

	logger.Info(ctx, "user retrieved successfully",
		zap.String("user_id", userID),
		zap.String("username", "john_doe"),
	)

	c.JSON(http.StatusOK, gin.H{
		"id":       userID,
		"username": "john_doe",
		"email":    "john@example.com",
	})
}

func CreateOrderHandler(c *gin.Context) {
	ctx := c.Request.Context()

	var order struct {
		ProductID string `json:"product_id"`
		Quantity  int    `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		logger.Warn(ctx, "invalid request body", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	logger.Info(ctx, "processing order",
		zap.String("product_id", order.ProductID),
		zap.Int("quantity", order.Quantity),
	)

	logger.Info(ctx, "validating stock",
		zap.String("product_id", order.ProductID),
	)
	time.Sleep(time.Duration(20+rand.Intn(30)) * time.Millisecond)

	logger.Info(ctx, "charging payment",
		zap.String("product_id", order.ProductID),
	)
	time.Sleep(time.Duration(40+rand.Intn(60)) * time.Millisecond)

	if rand.Float32() < 0.3 {
		err := errors.New("gateway timeout")
		logger.Error(ctx, "payment failed",
			zap.String("product_id", order.ProductID),
			zap.Error(err),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "payment failed"})
		return
	}

	logger.Info(ctx, "saving order to database",
		zap.String("product_id", order.ProductID),
	)
	time.Sleep(time.Duration(30+rand.Intn(40)) * time.Millisecond)

	orderID := "ORD-" + randomID()

	logger.Info(ctx, "order completed",
		zap.String("order_id", orderID),
		zap.String("product_id", order.ProductID),
		zap.Int("quantity", order.Quantity),
	)

	c.JSON(http.StatusCreated, gin.H{
		"order_id":   orderID,
		"product_id": order.ProductID,
		"quantity":   order.Quantity,
		"status":     "created",
	})
}

func GetProductsHandler(c *gin.Context) {
	ctx := c.Request.Context()

	logger.Info(ctx, "fetching product catalog")

	time.Sleep(time.Duration(50+rand.Intn(50)) * time.Millisecond)

	logger.Info(ctx, "catalog retrieved",
		zap.Int("product_count", 3),
	)

	c.JSON(http.StatusOK, gin.H{
		"products": []gin.H{
			{"id": "PROD-1", "name": "Product A", "price": 10.00},
			{"id": "PROD-2", "name": "Product B", "price": 20.00},
			{"id": "PROD-3", "name": "Product C", "price": 30.00},
		},
	})
}

func randomID() string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 6)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
