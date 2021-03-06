package main

import (
	"github.com/esma-yigit/partial-ecommerce/controllers"
	"github.com/esma-yigit/partial-ecommerce/database"
	"github.com/esma-yigit/partial-ecommerce/middleware"
	"github.com/esma-yigit/partial-ecommerce/routes"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	app := controllers.NewApplication(database.ProductData(database.Client, "Products"), database.UserData(database.Client, "Users"))
	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/addtocart", app.AddToCart())
	router.GET("/removeitem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	log.Fatal(router.Run(":", port))

}
