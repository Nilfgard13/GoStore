package controller

import (
	"net/http"

	"github.com/Nilfgard13/GOSTORE/app/model"
	"github.com/unrolled/render"
)

func (server *Server) Product(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout: "layout",
	})

	productModel := model.Product{}
	product, err := productModel.GetProduct(server.DB)
	if err != nil {
		return
	}

	_ = render.HTML(w, http.StatusOK, "product", map[string]interface{}{
		"product": product,
	})
}
