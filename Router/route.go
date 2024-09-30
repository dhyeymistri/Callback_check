package router

import (
	"github.com/gorilla/mux"

	SubscriptionController "Callback_check/Controllers/Subscriptions"
)

func GetRouter() *mux.Router {
	apiRouter := mux.NewRouter()
	newSubscriptionCallbackHandler(apiRouter)
	return apiRouter
}

func newSubscriptionCallbackHandler(router *mux.Router) {
	router.HandleFunc("/newSubscriptionMade", SubscriptionController.NotifyNewSub).Methods("POST")
}
