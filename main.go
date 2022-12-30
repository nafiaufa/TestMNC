package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Customer represents a customer in the system
type Customer struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// History represents a transaction history
type History struct {
	ID        int       `json:"id"`
	Customer  int       `json:"customer"`
	Timestamp time.Time `json:"timestamp"`
	Amount    int       `json:"amount"`
}

// Customers is a map containing all registered customers
var Customers = make(map[int]Customer)

// Histories is a slice containing all transaction histories
var Histories []History

// LoginHandler handles login requests
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Unmarshal request body into Customer struct
	var customer Customer
	err = json.Unmarshal(body, &customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if customer exists in Customers map
	_, exists := Customers[customer.ID]
	if !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Check if password is correct
	if Customers[customer.ID].Password != customer.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Login successful, write response
	response, err := json.Marshal(customer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

// PaymentHandler handles payment requests
func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Unmarshal request body into History struct
	var history History
	err = json.Unmarshal(body, &history)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if customer is logged in
	if _, exists := Customers[history.Customer]; !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Perform payment
	history.Timestamp = time.Now()
	Histories = append(Histories, history)

	// Write response
	response, err := json.Marshal(history)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

// LogoutHandler handles logout requests
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Unmarshal request body into Customer struct
	var customer Customer
	err = json.Unmarshal(body, &customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if customer is logged in
	if _, exists := Customers[customer.ID]; !exists {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Logout customer
	delete(Customers, customer.ID)

	// Write response
	response, err := json.Marshal(customer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

func main() {
	// Load customer data from JSON file
	customersJSON, err := ioutil.ReadFile("customers.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(customersJSON, &Customers)
	if err != nil {
		log.Fatal(err)
	}
	// Set up HTTP handlers
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/payment", PaymentHandler)
	http.HandleFunc("/logout", LogoutHandler)

	// Start HTTP server
	fmt.Println("Starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
