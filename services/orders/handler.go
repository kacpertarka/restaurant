package orders

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kacpertarka/restaurant/utils"
)

type OrderHandler struct {
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

func (handler *OrderHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/healthcheck", handler.healthChecker).Methods("GET")
}

func (handler *OrderHandler) healthChecker(w http.ResponseWriter, r *http.Request) {
	// test route

	utils.WriteJSON(w, http.StatusOK, fmt.Sprint("Your health is OK"))
}
