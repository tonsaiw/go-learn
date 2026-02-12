package handler

import (
	"fmt"
	"net/http"
)

type Order struct {}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Order created")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List all orders")
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get order by ID")
}

func (o *Order) Update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update order")
}

func (o *Order) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete order")
}