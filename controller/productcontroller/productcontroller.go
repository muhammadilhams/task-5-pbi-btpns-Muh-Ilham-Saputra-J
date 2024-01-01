package productcontroller

import (
	"net/http"

	"github.com/jeypc/go-jwt-mux/helper"
)

func Index(w http.ResponseWriter, r *http.Request ){

	data := []map[string]interface{}{
		{
			"id" : 1,
			"nama_product": "kemeja",
			"stok": 1000,

		},
		{
			"id" : 2,
			"nama_product": "celana",
			"stok": 2000,

		},
		{
			"id" : 1,
			"nama_product": "Sepatu",
			"stok": 3000,

		},
	}
	helper.ResponseJSON(w, http.StatusOK, data)
}
