package main

import (
	"GoTask/crud"
	"GoTask/transactions"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)



func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", crud.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", crud.GetUser).Methods("GET")
	r.HandleFunc("/users", crud.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", crud.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", crud.DeleteUser).Methods("DELETE")
	r.HandleFunc("/accounts/{id_user}", transactions.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{id}", transactions.DeleteAccount).Methods("DELETE")
	r.HandleFunc("/accounts/{id}", transactions.GetBalance).Methods("GET")
	r.HandleFunc("/accounts/{acc_id}", transactions.MoneyTransfer).Methods("PUT")
	r.HandleFunc("/accountsdepositMoney", transactions.DepositMoney).Methods("PUT")
	r.HandleFunc("/accountswithdrawal", transactions.WithdrawalMoney).Methods("PUT")
	r.HandleFunc("/accountstransactions/{id_acc_sender}", transactions.GetTransactions).Methods("PUT")
	r.HandleFunc("/transactions/{id}", transactions.CancelTransaction).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}
