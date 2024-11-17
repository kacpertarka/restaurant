package orders

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kacpertarka/restaurant/utils"
)

type OrderHandler struct {
	db *sql.DB
}

func NewOrderHandler(db *sql.DB) *OrderHandler {
	return &OrderHandler{db: db}
}

func (handler *OrderHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/healthcheck", handler.healthChecker).Methods("GET")
}

func (handler *OrderHandler) healthChecker(w http.ResponseWriter, r *http.Request) {
	// test route

	utils.WriteJSON(w, http.StatusOK, fmt.Sprint("Your health is OK!!!"))
}
