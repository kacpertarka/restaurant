package orders

import (
	"net/http"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{}
}

func (handler *OrderHandler) RegisterRoutes(router mux.Router) {
	router.HandleFunc("/healthcheck", handler.healthChecker)
}

func (handler *OrderHandler) healthChecker(w http.ResponseWriter, r *http.Request) {
	// test route

}
