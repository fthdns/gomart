package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fthdns/gomart/app/models"
	"github.com/unrolled/render"
)

func (server *Server) Products(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout: "layout",
	})

	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	if page <= 0 {
		page = 1
	}

	perPage := 9

	productModel := models.Product{}
	products, totalRows, err := productModel.GetProducts(server.DB, perPage, page)
	if err != nil {
		return
	}

	fmt.Println("====", totalRows)

	_ = render.HTML(w, http.StatusOK, "products", map[string]interface{}{
		"products": products,
	})
}
