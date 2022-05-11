package routes

import (
	"github.com/gorilla/mux"
	"github.com/morozvol/money_manager_api/internal/controllers"
	"github.com/morozvol/money_manager_api/pkg/store"
)

func AccountRoute(router *mux.Router, store store.Store) {
	router.HandleFunc("/account", controllers.CreateAccount(store)).Methods("POST")
	router.HandleFunc("/account/{accountId}", controllers.DeleteAccount(store)).Methods("DELETE")
	router.HandleFunc("/account/{accountId}", controllers.EditAccount(store)).Methods("PUT")
	router.HandleFunc("/account/{accountId}", controllers.GetAccount(store)).Methods("GET")
	router.HandleFunc("/accounts/{userId}", controllers.GetAccounts(store)).Methods("GET")
}
