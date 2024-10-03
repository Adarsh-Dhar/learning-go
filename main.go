package main

import (
	"log"
	"os"

	"github.com/Adarsh-Dhar/learning-go/controllers"
	"github.com/Adarsh-Dhar/learning-go/database"
	"github.com/Adarsh-Dhar/learning-go/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))

	router := gin.New()
	router.Use(gin.Logger())

	router.UserRoutes(router)
	router.Use(middlewares.Authentication())

	router.Get("/addtocart", app.AddToCart())
	router.Get("/removeitem", app.RemoveItem())
	router.Get("/cartcheckout", app.BuyFromCart())
	router.Get("/instantbuy", app.InstaBuy())

	log.Fatal(router.Run(":" + port))
}
