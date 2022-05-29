package controllers

import (
	"context"
	"errors"
	"github.com/esma-yigit/partial-ecommerce/database"
	"github.com/esma-yigit/partial-ecommerce/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

func (app *Application) AddToCart() gin.Handler {
	return func(c *gin.Context) {
		productQueryId := c.Query("id")
		if productQueryId == "" {
			log.Panicln("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}
		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Panicln("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("use id is empty"))
			return
		}
		productID, err := primitive.ObjectIDFromHex(productQueryId)
		if err != nil {
			log.Panicln(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := database.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(200, "Successfully added to cart")
		}
	}

}
func (app *Application) RemoveItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryId := c.Query("id")
		if productQueryId == "" {
			log.Panicln("product id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}
		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Panicln("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("use id is empty"))
			return
		}
		productID, err := primitive.ObjectIDFromHex(productQueryId)
		if err != nil {
			log.Panicln(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := database.RemoveCartItem(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Successfully removed item from cart")
	}
}
func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryId := c.Query("id")
		if productQueryId == "" {
			log.Panicln("product id is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}``
		userQueryID := c.Query("userID")
		if userQueryID == "" {
			log.Panicln("user id is empty")
			_ = c.AbortWithError(http.StatusBadRequest, errors.New("use id is empty"))
			return
		}
		productID, err := primitive.ObjectIDFromHex(productQueryId)
		if err != nil {
			log.Panicln(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err := database.InstantBuyer(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		c.IndentedJSON(200, "Successfully placed the order")
	}
}
func (app *Application) BuyFromCart() gin.HandlerFunc {
return func(c *gin.Context) {
	userQueryID := c.Query("userID")
	if userQueryID == "" {
		log.Panicln("user id is empty")
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("use id is empty"))
		return
	}
	var ctx,cancel:=context.WithTimeout(context.Background(),100*time.Second)
	err := database.BuyItemFromCart(ctx, app.userCollection, userQueryID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(200, "Successfully placed the order")
}
}


func GetItemFromCart() gin.HandlerFunc {
return func(c *gin.Context) {
	user_id:=c.Query("id")
	if user_id=="" {
		c.Header("Content-Type","application/json")
		c.JSON(http.StatusNotFound,gin.H{"Error":"Invalid id"})
		c.Abort()
		return
	}
	usert_id,_:=primitive.ObjectIDFromHex(user_id)
	var ctx,cancel=context.WithTimeout(context.Background(),100*time.Second)
	defer cancel()
	var filledcart models.User
	err:=UserCollection.FindOne(ctx,bson.D{primitive.E{Key: "_id",Value: usert_id}}).Decode(&filledcart)
	if err!=nil{
		log.Println(err)
		c.IndentedJSON(500,"not found")
	}
	filter_match:=bson.D{{Key: "$match",Value: bson.D{primitive.E{Key: "_id",Value: usert_id}}}}
	unwind:=bson.D{{Key: "$unwind",Value: bson.D{primitive.E{Key:"path",Value: "$usercart"}}}}
	groping:=bson.D{{Key: "$group",Value: bson.D{primitive.E{Key: "_id",Value: "$_id"},{Key: "total",Value: primitive.D{primitive.E{Key:"$sum",Value: "$usercart.price"}}}}}}
    pointcursor,err := UserCollection.Aggregate(ctx,mongo.Pipeline{filter_match,unwind,groping})
	if err != nil {
		log.Println(err)

	}
	var listing []bson.M
	if err=pointcursor.All(ctx,&listing);err!=nil {
		log.Panicln(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	for _, json := range listing {
		c.IndentedJSON(200,json["total"])
		c.IndentedJSON(200,filledcart.UserCart)
	}
	ctx.Done()


}
}
