package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name      *string            `json:"first_name"`
	Last_Name       *string            `json:"last_name"`
	Password        *string            `json:"password"`
	Email           *string            `json:"email"`
	Phone           *string            `json:"phone"`
	Token           *string            `json:"token"`
	Refresh_Token   *string            `json:"refresh_token"`
	Created_At      time.Time          `json:"created_at"`
	Updated_At      time.Time          `json:"updated_at"`
	User_ID         string             `json:"user_id"`
	UserCart        []ProductUser      `json:"usercart" bson:"usercart"`
	Address_Details []Address          `json:"address" bson:"address"`
	Order_Status    []Order            `json:"orders" bson:"orders"`
}
type Product struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name" `
	Price        *uint              `json:"price"`
	Rating       *uint8             `json:"rating"`
	image        *string            `json:"image"`
}
type ProductUser struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string            `json:"product_name" bson:"product_name"`
	Price        *uint              `json:"price" bson:"price"`
	Rating       *uint              `json:"rating" bson:"rating"`
	image        *string            `json:"image" json:"image"`
}
type Address struct {
	Address_ID primitive.ObjectID `bson:"_id"`
	House      *string            `json:"house_name" bson:"house_name"`
	Street     *string            `json:"street_name" bson:"street_name"`
	City       *string            `json:"city_name" bson:"city_name"`
	Pincode    *string            `json:"pin_code" bson:"pin_code"`
}
type Order struct {
	Order_ID        primitive.ObjectID `bson:"_id"`
	Order_Cart      []ProductUser      `json:"order_list" bson:"order_list"`
	Ordered_At      time.Time          `json:"ordered_at" bson:"ordered_at"`
	Price           *uint              `json:"total_price" bson:"total_price"`
	Discount        *int               `json:"discount" bson:"discount"`
	Payment_Methods Payment            `json:"payment_method" bson:"payment_method"`
}
type Payment struct {
	Digital bool
	COD     bool
}
