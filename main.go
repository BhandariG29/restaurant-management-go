package main

import (
	"os"

	database "github.com/BhandariG29/restaurant-management-go/database"
	middleware "github.com/BhandariG29/restaurant-management-go/middleware"
	routes "github.com/BhandariG29/restaurant-management-go/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main(){
	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.InvoiceRoutes(router)
	routes.MenuRoutes(router)
	routes.OrderItemRoutes(router)
	routes.OrderRoutes(router)
	routes.TableRoutes(router)

	router.Run(":" + port)


}