package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("Cant find the product")
	ErrCantDecodeProducts = errors.New("Cant find the products")
	ErrUserIdInvalid      = errors.New("This user is invalid")
	ErrCantUpdateUser     = errors.New("Cant")
	ErrCantRemoveItemCart = errors.New("cannot remove item from the cart")
	ErrCantGetItem        = errors.New("was unable to get the item from the cart")
	ErrCantBuyCartItem    = errors.New("cannot update the purchase")
)

func AddProductToCart() {

}
func RemoveCartItem() {

}
func BuyItemFromCart() {

}
func InstantBuyer() {

}
