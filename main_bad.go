package main

// import (
// 	"log"
// 	"math/rand"
// 	"net/http"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	log.Println("Starting server...")

// 	gin.SetMode(gin.ReleaseMode)
// 	r := gin.New()

// 	r.GET("/users/:id", func(c *gin.Context) {
// 		log.Println("Fetching user from database")

// 		time.Sleep(time.Duration(30+rand.Intn(50)) * time.Millisecond)

// 		// Error aleatorio - no sabes de qué request viene
// 		if rand.Float32() < 0.3 {
// 			log.Println("ERROR: Database connection timeout")
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
// 			return
// 		}

// 		log.Println("User retrieved successfully")
// 		c.JSON(http.StatusOK, gin.H{"id": c.Param("id"), "username": "john_doe"})
// 	})

// 	r.POST("/orders", func(c *gin.Context) {
// 		var order struct {
// 			ProductID string `json:"product_id"`
// 			Quantity  int    `json:"quantity"`
// 		}
// 		c.ShouldBindJSON(&order)

// 		log.Println("Processing order")

// 		log.Println("Validating stock")
// 		time.Sleep(time.Duration(20+rand.Intn(30)) * time.Millisecond)

// 		log.Println("Charging payment")
// 		time.Sleep(time.Duration(40+rand.Intn(60)) * time.Millisecond)

// 		// Error aleatorio de pago
// 		if rand.Float32() < 0.3 {
// 			log.Println("ERROR: Payment gateway timeout")
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "payment failed"})
// 			return
// 		}

// 		log.Println("Saving order to database")
// 		time.Sleep(time.Duration(30+rand.Intn(40)) * time.Millisecond)

// 		log.Println("Order completed")
// 		c.JSON(http.StatusCreated, gin.H{"order_id": "ORD-12345"})
// 	})

// 	r.GET("/products", func(c *gin.Context) {
// 		log.Println("Fetching product catalog")
// 		time.Sleep(time.Duration(50+rand.Intn(50)) * time.Millisecond)
// 		log.Println("Catalog retrieved")
// 		c.JSON(http.StatusOK, gin.H{"products": []string{"A", "B", "C"}})
// 	})

// 	log.Println("Server running on :8080")
// 	r.Run(":8080")
// }

// // Problemas:
// // 1. Sin trace-id → imposible correlacionar logs
// // 2. Logs con fmt/log → sin estructura, difícil parsear
// // 3. Sin contexto → no sabemos de qué request viene cada log
// // 4. Sin niveles
