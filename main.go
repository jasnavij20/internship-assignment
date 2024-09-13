package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// Account represents the structure of an account asset
type Account struct {
	DealerID    string  `json:"dealer_id"`
	MSISDN      string  `json:"msisdn"`
	MPIN        string  `json:"mpin"`
	Balance     float64 `json:"balance"`
	Status      string  `json:"status"`
	TransAmount float64 `json:"trans_amount"`
	TransType   string  `json:"trans_type"`
	Remarks     string  `json:"remarks"`
}

// assetDB simulates a database using a map with concurrency control
var assetDB = make(map[string]Account)
var dbMutex sync.Mutex

// AddAccount handles adding a new account to the system
func AddAccount(w http.ResponseWriter, r *http.Request) {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	var newAccount Account
	if err := json.NewDecoder(r.Body).Decode(&newAccount); err != nil {
		http.Error(w, "Invalid data format", http.StatusBadRequest)
		return
	}

	if _, exists := assetDB[newAccount.DealerID]; exists {
		http.Error(w, "Account already exists", http.StatusConflict)
		return
	}

	assetDB[newAccount.DealerID] = newAccount
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newAccount)
}

// GetAccount retrieves an account by DealerID
func GetAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dealerID := params["dealerId"]

	dbMutex.Lock()
	defer dbMutex.Unlock()

	account, found := assetDB[dealerID]
	if !found {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(account)
}

// UpdateAccount modifies an existing account
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dealerID := params["dealerId"]

	dbMutex.Lock()
	defer dbMutex.Unlock()

	var updatedAccount Account
	if err := json.NewDecoder(r.Body).Decode(&updatedAccount); err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	if _, exists := assetDB[dealerID]; !exists {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	assetDB[dealerID] = updatedAccount
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedAccount)
}

// DeleteAccount removes an account from the system
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dealerID := params["dealerId"]

	dbMutex.Lock()
	defer dbMutex.Unlock()

	if _, found := assetDB[dealerID]; !found {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	delete(assetDB, dealerID)
	w.WriteHeader(http.StatusNoContent)
}

// ListAllAccounts returns all accounts currently stored
func ListAllAccounts(w http.ResponseWriter, r *http.Request) {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	var accounts []Account
	for _, account := range assetDB {
		accounts = append(accounts, account)
	}

	json.NewEncoder(w).Encode(accounts)
}

// MockTransactionHistory returns a mock transaction history for an account
func MockTransactionHistory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dealerID := params["dealerId"]

	dbMutex.Lock()
	defer dbMutex.Unlock()

	account, found := assetDB[dealerID]
	if !found {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}

	// This is a mock; in a real system, you'd pull from actual historical data
	history := []Account{account}
	json.NewEncoder(w).Encode(history)
}

func main() {
	// Initialize the router
	router := mux.NewRouter()

	// Route handlers for managing accounts
	router.HandleFunc("/accounts", AddAccount).Methods("POST")
	router.HandleFunc("/accounts", ListAllAccounts).Methods("GET")
	router.HandleFunc("/accounts/{dealerId}", GetAccount).Methods("GET")
	router.HandleFunc("/accounts/{dealerId}", UpdateAccount).Methods("PUT")
	router.HandleFunc("/accounts/{dealerId}", DeleteAccount).Methods("DELETE")
	router.HandleFunc("/accounts/{dealerId}/history", MockTransactionHistory).Methods("GET")

	// Start the server
	fmt.Println("API server running at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

