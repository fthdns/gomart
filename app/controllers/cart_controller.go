package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"uuid"

	"github.com/fthdns/gomart/app/models"
	"gorm.io/gorm"
)

func GetShoppingCartID(w http.ResponseWriter, r *http.Request) string {
	session, _ := store.Get(r, "shopping-cart-session")

	if session.Values["cart-id"] == nil {
		session.Values["cart-id"] = uuid.New().String()
		session.Save(r, w)
	}

	return fmt.Sprint("%s", session.Values["cart-id"])
}

func GetShoppingCart(db *gorm.DB, cartID string) (*models.Cart, error) {
	var cart models.Cart

	existCart, err := cart.GetCart(db, cartID)

	if err != nil {
		existCart, _ = cart.CreateCart(db, cartID)
	}

	return existCart, nil
}

func (server *Server) GetCart(w http.ResponseWriter, r *http.Request) {
	panic("cart items")
}

func (server *Server) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	productID := r.FormValue("product_id")
	qty, _ := strconv.Atoi(r.FormValue("qty"))

	productModel := models.Product{}
	product, err := productModel.FindByID(server.DB, productID)

	if err != nil {
		http.Redirect(w, r, "/products/"+product.Slug, http.StatusSeeOther)
	}

	if qty > product.Stock {
		http.Redirect(w, r, "/products/"+product.Slug, http.StatusSeeOther)
	}

	var cart *models.Cart

	cartID := GetShoppingCartID(w, r)
	cart, _ = GetShoppingCart(server.DB, cartID)

	http.Redirect(w, r, "/carts", http.StatusSeeOther)
}
