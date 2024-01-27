package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Nilfgard13/GOSTORE/app/model"
	"github.com/unrolled/render"
)

func (server *Server) Product(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout: "layout",
	})

	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	if page <= 0 {
		page = 1
	}

	perPage := 9

	productModel := model.Product{}
	product, totalRow, err := productModel.GetProduct(server.DB, perPage, page)
	if err != nil {
		return
	}

	pagination, _ := GetPaginationLink(server.AppConfig, PaginationParams{
		Path:        "product",
		TotalRow:    int32(totalRow),
		PerPage:     int32(perPage),
		CurrentPage: int32(page),
	})

	fmt.Println("===", pagination)

	_ = render.HTML(w, http.StatusOK, "product", map[string]interface{}{
		"product":    product,
		"pagination": pagination,
	})
}
