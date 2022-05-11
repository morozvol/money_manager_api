package controllers

import (
	"github.com/morozvol/money_manager_api/pkg/store"
	"github.com/morozvol/money_manager_api/tools"
	"net/http"
)

func GetCategoryTree(store store.Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		tree, err := store.Category().Get(1000000)
		if err != nil {
			tools.SendError(rw, http.StatusInternalServerError, err)
			return
		}
		tools.SendResponse(rw, http.StatusOK, "success", *tree)
	}
}
