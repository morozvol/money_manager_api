package routes

import (
	"github.com/gorilla/mux"
	"github.com/morozvol/money_manager_api/internal/controllers"
	"github.com/morozvol/money_manager_api/pkg/store"
)

func CategoryRoute(router *mux.Router, store store.Store) {
	//router.HandleFunc("/operation", controllers.CreateOperation(store)).Methods("POST")
	//router.HandleFunc("/operation/{operationId}", controllers.GetOperation(store)).Methods("GET")
	//router.HandleFunc("/operation/{operationId}", controllers.EditAUser()).Methods("PUT")
	//router.HandleFunc("/operation/{operationId}", controllers.DeleteAUser()).Methods("DELETE")
	router.HandleFunc("/categories", controllers.GetCategoryTree(store)).Methods("GET")
}
