package main

import (
	"os"
	"restaurant-api/database"
	"restaurant-api/middleware"
	"restaurant-api/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItem(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
