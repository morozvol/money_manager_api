package routes

import (
	"github.com/gorilla/mux"
	"github.com/morozvol/money_manager_api/internal/controllers"
	"github.com/morozvol/money_manager_api/pkg/store"
)

func OperationRoute(router *mux.Router, store store.Store) {
	router.HandleFunc("/operation", controllers.CreateOperation(store)).Methods("POST")
	router.HandleFunc("/operation/{operationId}", controllers.GetOperation(store)).Methods("GET")
	router.HandleFunc("/operation/{operationId}", controllers.EditOperation(store)).Methods("PUT")
	router.HandleFunc("/operation/{operationId}", controllers.DeleteOperation(store)).Methods("DELETE")
	//router.HandleFunc("/operations/{dateFrom}/{dateTo}", controllers.GetOperation(store)).Methods("GET")

}
