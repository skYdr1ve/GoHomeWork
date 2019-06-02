package transactions

import (
	"GoTask/settings"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

const connStr = "user=postgres password=alemip10 dbname=BankSystem sslmode=disable"

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	log.Println("Function CreateAccount")
	w.Header().Set("Content-Type", "application/json")
	var db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println("There was an error connecting to the database")
		defer db.Close()
		err = json.NewEncoder(w).Encode("500 Internal Server Error")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}else{
		log.Println("Database connection established")
	}
	defer db.Close()
	params := mux.Vars(r)
	var acc settings.Account
	err = json.NewDecoder(r.Body).Decode(&acc)
	if err != nil {
		log.Println("Could not get request body")
		err = json.NewEncoder(w).Encode("500 Internal Server Error")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}else{
		log.Println("The request body was received")
	}
	_, err = db.Exec("insert into accounts (id_user, amount) values ($1, $2)", params["id_user"], acc.Amount)
	if err != nil {
		log.Println("Could not execute query to DB")
		err = json.NewEncoder(w).Encode("500 Internal Server Error")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}else{
		log.Println("The request has been completed")
	}
	err = json.NewEncoder(w).Encode("Account has been created")
	if err != nil {
		log.Panic("Could not return responce")
	}
	log.Println("The function was successful")
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	log.Println("Function DeleteAccount")
	w.Header().Set("Content-Type", "application/json")
	var db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println("There was an error connecting to the database")
		defer db.Close()
		log.Println("Could not execute query to DB")
		err = json.NewEncoder(w).Encode("500 Internal Server Error")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}else{
		log.Println("Database connection established")
	}
	defer db.Close()
	params := mux.Vars(r)
	tx, err := db.Begin()
	if err != nil {
		log.Println("Could not create transactions")
		err = json.NewEncoder(w).Encode("500 Internal Server Error")
		if err != nil {
			log.Panic("Could not return responce")
		}
	}
	{
		stmt, err := tx.Prepare("delete from transactions where id_acc_sender = $1")
		if err != nil {
			tx.Rollback()
			log.Println("Could not execute query to DB")
			defer stmt.Close()
			err = json.NewEncoder(w).Encode("500 Internal Server Error")
			if err != nil {
				log.Panic("Could not return responce")
			}
			return
		}
		defer stmt.Close()
		_, err = stmt.Exec(params["id"])
		if err != nil {
			tx.Rollback()
			log.Println("Could not execute query to DB")
			err = json.NewEncoder(w).Encode("500 Internal Server Error")
			if err != nil {
				log.Panic("Could not return responce")
			}
			return
		}
	}
	{
		stmt, err := tx.Prepare("delete from accounts where id = $1")
		if err != nil {
			tx.Rollback()
			log.Println("Could not execute query to DB")
			defer stmt.Close()
			err = json.NewEncoder(w).Encode("500 Internal Server Error")
			if err != nil {
				log.Panic("Could not return responce")
			}
			return
		}
		defer stmt.Close()
		_, err = stmt.Exec(params["id"])
		if err != nil {
			tx.Rollback()
			log.Println("Could not execute query to DB")
			err = json.NewEncoder(w).Encode("500 Internal Server Error")
			if err != nil {
				log.Panic("Could not return responce")
			}
			return
		}
	}
	tx.Commit()

	err = json.NewEncoder(w).Encode("Account has been deleted")
	if err != nil {
		log.Panic("Could not return responce")
	}
	log.Println("The function was successful")
}

func GetBalance(w http.ResponseWriter, r *http.Request) {
	log.Println("Function GetBalance")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println("There was an error connecting to the database")
		defer db.Close()
		err = json.NewEncoder(w).Encode("500 Internal Server Error")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}else{
		log.Println("Database connection established")
	}
	defer db.Close()
	row := db.QueryRow("select * from accounts where id = $1", params["id"])
	account := settings.Account{}
	err = row.Scan(&account.Id, &account.Id_user, &account.Amount)
	if err != nil {
		err = json.NewEncoder(w).Encode("No such account")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	} else {
		err = json.NewEncoder(w).Encode(account.Amount)
		if err != nil {
			log.Panic("Could not return responce")
		}
	}
	log.Println("The function was successful")
}

func MoneyTransfer(w http.ResponseWriter, r *http.Request) {
	log.Println("Function MoneyTransfer")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println("There was an error connecting to the database")
		defer db.Close()
		err = json.NewEncoder(w).Encode("500 Internal Server Error")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}else{
		log.Println("Database connection established")
	}
	defer db.Close()

	account1 := settings.Account{}
	account2 := settings.Account{}
	err = json.NewDecoder(r.Body).Decode(&account1)
	if err != nil {
		log.Println("Could not get request body")
		err = json.NewEncoder(w).Encode("500 Internal Server Error")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}else{
		log.Println("The request body was received")
	}
	row := db.QueryRow("select * from accounts where id = $1", params["acc_id"])
	err = row.Scan(&account2.Id, &account2.Id_user, &account2.Amount)
	if err != nil {
		err = json.NewEncoder(w).Encode("No such account")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}
	amount_2,err := strconv.Atoi(account2.Amount)
	if err != nil {
		err = json.NewEncoder(w).Encode("500 Internal Server Error")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}
	amount_1,err := strconv.Atoi(account1.Amount)
	if err != nil {
		err = json.NewEncoder(w).Encode("500 Internal Server Error")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}
	if amount_2 < amount_1 {
		err = json.NewEncoder(w).Encode("You do not have enough money")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}
	tx, err := db.Begin()
	if err != nil {
		log.Println("Could not create transactions")
		err = json.NewEncoder(w).Encode("500 Internal Server Error")
		if err != nil {
			log.Panic("Could not return responce")
		}
	}
	{
		stmt, err := tx.Prepare("update accounts set amount=amount - $1 where id = $2;")
		if err != nil {
			tx.Rollback()
			log.Println("Could not execute query to DB")
			defer stmt.Close()
			err = json.NewEncoder(w).Encode("500 Internal Server Error")
			if err != nil {
				log.Panic("Could not return responce")
			}
			return
		}
		defer stmt.Close()
		_, err = stmt.Exec(account1.Amount, account2.Id)
		if err != nil {
			tx.Rollback()
			log.Println("Could not execute query to DB")
			defer stmt.Close()
			err = json.NewEncoder(w).Encode("500 Internal Server Error")
			if err != nil {
				log.Panic("Could not return responce")
			}
			return
		}
	}
	{
		stmt, err := tx.Prepare("update accounts set amount=amount + $1 where id = $2;")
		if err != nil {
			tx.Rollback()
			log.Println("Could not execute query to DB")
			defer stmt.Close()
			err = json.NewEncoder(w).Encode("500 Internal Server Error")
			if err != nil {
				log.Panic("Could not return responce")
			}
			return
		}
		defer stmt.Close()
		_, err = stmt.Exec(account1.Amount, account1.Id)
		if err != nil {
			tx.Rollback()
			log.Println("Could not execute query to DB")
			defer stmt.Close()
			err = json.NewEncoder(w).Encode("500 Internal Server Error")
			if err != nil {
				log.Panic("Could not return responce")
			}
			return
		}
	}
	{
		stmt, err := tx.Prepare("insert into transactions (id_acc_reciver,id_acc_sender,type,date,amount) values ($1,$2,$3,$4,$5);")
		if err != nil {
			tx.Rollback()
			log.Println("Could not execute query to DB")
			defer stmt.Close()
			err = json.NewEncoder(w).Encode("500 Internal Server Error")
			if err != nil {
				log.Panic("Could not return responce")
			}
			return
		}
		defer stmt.Close()
		_, err = stmt.Exec(account1.Id, account2.Id, "transfer", time.Now(), account1.Amount)
		if err != nil {
			tx.Rollback()
			log.Println("Could not execute query to DB")
			defer stmt.Close()
			err = json.NewEncoder(w).Encode("500 Internal Server Error")
			if err != nil {
				log.Panic("Could not return responce")
			}
			return
		}
	}
	tx.Commit()
	/*db.Exec("update accounts set amount=amount - $1 where id = $2", account1.Amount, account2.Id)
	db.Exec("update accounts set amount=amount + $1 where id = $2", account1.Amount, account1.Id)
	db.Exec("insert into transactions (id_acc_reciver,id_acc_sender,type,date,amount) values ($1,$2,$3,$4,$5)",
		account1.Id, account2.Id, "transfer", time.Now(), account1.Amount)*/
	err = json.NewEncoder(w).Encode("Money transfer was successful")
	if err != nil {
		log.Panic("Could not return responce")
	}
	log.Println("The function was successful")
}

func DepositMoney(w http.ResponseWriter, r *http.Request) {
	log.Println("Function DepositMoney")
	w.Header().Set("Content-Type", "application/json")
	var db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println("There was an error connecting to the database")
		defer db.Close()
		err = json.NewEncoder(w).Encode("500 Internal Server Error")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}else{
		log.Println("Database connection established")
	}
	defer db.Close()
	account := settings.Account{}
	err = json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		log.Println("Could not get request body")
		err = json.NewEncoder(w).Encode("500 Internal Server Error")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}else{
		log.Println("The request body was received")
	}
	tx, err := db.Begin()
	if err != nil {
		log.Println("Could not create transactions")
		err = json.NewEncoder(w).Encode("500 Internal Server Error")
		if err != nil {
			log.Panic("Could not return responce")
		}
	}
	{
		stmt, err := tx.Prepare("update accounts set amount=amount + $1 where id = $2;")
		if err != nil {
			tx.Rollback()
			log.Println("Could not execute query to DB")
			defer stmt.Close()
			err = json.NewEncoder(w).Encode("500 Internal Server Error")
			if err != nil {
				log.Panic("Could not return responce")
			}
			return
		}
		defer stmt.Close()
		_, err = stmt.Exec(account.Amount, account.Id)
		if err != nil {
			tx.Rollback()
			log.Println("Could not execute query to DB")
			defer stmt.Close()
			err = json.NewEncoder(w).Encode("500 Internal Server Error")
			if err != nil {
				log.Panic("Could not return responce")
			}
			return
		}
	}
	{
		stmt, err := tx.Prepare("insert into transactions (id_acc_reciver,id_acc_sender,type,date,amount) values ($1,$2,$3,$4,$5);")
		if err != nil {
			tx.Rollback()
			log.Println("Could not execute query to DB")
			defer stmt.Close()
			err = json.NewEncoder(w).Encode("500 Internal Server Error")
			if err != nil {
				log.Panic("Could not return responce")
			}
			return
		}
		defer stmt.Close()
		_, err = stmt.Exec(account.Id, account.Id, "deposit", time.Now(), account.Amount)
		if err != nil {
			tx.Rollback()
			log.Println("Could not execute query to DB")
			defer stmt.Close()
			err = json.NewEncoder(w).Encode("500 Internal Server Error")
			if err != nil {
				log.Panic("Could not return responce")
			}
			return
		}
	}
	tx.Commit()
	err = json.NewEncoder(w).Encode("The money was successfully credited")
	if err != nil {
		log.Panic("Could not return responce")
	}
	log.Println("The function was successful")
}

func WithdrawalMoney(w http.ResponseWriter, r *http.Request) {
	log.Println("Function WithdrawalMoney")
	w.Header().Set("Content-Type", "application/json")
	var db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println("There was an error connecting to the database")
		defer db.Close()
		err = json.NewEncoder(w).Encode("500 Internal Server Error")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}else{
		log.Println("Database connection established")
	}
	defer db.Close()
	account1 := settings.Account{}
	account2 := settings.Account{}
	err = json.NewDecoder(r.Body).Decode(&account1)
	if err != nil {
		log.Println("Could not get request body")
		err = json.NewEncoder(w).Encode("500 Internal Server Error")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}else{
		log.Println("The request body was received")
	}
	row := db.QueryRow("select * from accounts where id = $1", account1.Id)
	err = row.Scan(&account2.Id, &account2.Id_user, &account2.Amount)
	if err != nil {
		err = json.NewEncoder(w).Encode("No such account")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	} else {
		if account1.Amount > account2.Amount {
			err = json.NewEncoder(w).Encode("Not enough money")
			if err != nil {
				log.Panic("Could not return responce")
			}
			return
		}else{
			tx, err := db.Begin()
			if err != nil {
				log.Println("Could not create transactions")
				err = json.NewEncoder(w).Encode("500 Internal Server Error")
				if err != nil {
					log.Panic("Could not return responce")
				}
			}
			{
				stmt, err := tx.Prepare("update accounts set amount=amount - $1 where id = $2;")
				if err != nil {
					tx.Rollback()
					log.Println("Could not execute query to DB")
					defer stmt.Close()
					err = json.NewEncoder(w).Encode("500 Internal Server Error")
					if err != nil {
						log.Panic("Could not return responce")
					}
					return
				}
				defer stmt.Close()
				_, err = stmt.Exec(account1.Amount, account1.Id)
				if err != nil {
					tx.Rollback()
					log.Println("Could not execute query to DB")
					defer stmt.Close()
					err = json.NewEncoder(w).Encode("500 Internal Server Error")
					if err != nil {
						log.Panic("Could not return responce")
					}
					return
				}
			}
			{
				stmt, err := tx.Prepare("insert into transactions (id_acc_reciver,id_acc_sender,type,date,amount) values ($1,$2,$3,$4,$5);")
				if err != nil {
					tx.Rollback()
					log.Println("Could not execute query to DB")
					defer stmt.Close()
					err = json.NewEncoder(w).Encode("500 Internal Server Error")
					if err != nil {
						log.Panic("Could not return responce")
					}
					return
				}
				defer stmt.Close()
				_, err = stmt.Exec(account1.Id, account1.Id, "withdrawal", time.Now(), account1.Amount)
				if err != nil {
					tx.Rollback()
					log.Println("Could not execute query to DB")
					defer stmt.Close()
					err = json.NewEncoder(w).Encode("500 Internal Server Error")
					if err != nil {
						log.Panic("Could not return responce")
					}
					return
				}
			}
			tx.Commit()
			err = json.NewEncoder(w).Encode("Money has been successfully withdrawn")
			if err != nil {
				log.Panic("Could not return responce")
			}
		}
	}
	log.Println("The function was successful")
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	log.Println("Function GetTransactions")
	w.Header().Set("Content-Type", "application/json")
	var db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println("There was an error connecting to the database")
		defer db.Close()
		err = json.NewEncoder(w).Encode("500 Internal Server Error")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}else{
		log.Println("Database connection established")
	}
	defer db.Close()
	params := mux.Vars(r)
	date := settings.Date{}
	transactions := []settings.Transaction{}
	err = json.NewDecoder(r.Body).Decode(&date)
	if err != nil {
		log.Println("Could not get request body")
		err = json.NewEncoder(w).Encode("500 Internal Server Error")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}else{
		log.Println("The request body was received")
	}
	rows, err := db.Query("select * from transactions where id_acc_sender = $1 and date between $2 and $3",
		params["id_acc_sender"], date.Date_first, date.Date_second)
	if err != nil {
		defer rows.Close()
		err = json.NewEncoder(w).Encode("There are no such transactions")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}
	defer rows.Close()
	for rows.Next() {
		p := settings.Transaction{}
		err := rows.Scan(&p.Id, &p.Id_acc_reciver, &p.Id_acc_sender, &p.Type, &p.Date, &p.Amount)
		if err != nil {
			log.Println(err)
			continue
		}
		transactions = append(transactions, p)
	}
	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		log.Panic("Could not return responce")
	}
	log.Println("The function was successful")
}

func CancelTransaction(w http.ResponseWriter, r *http.Request) {
	log.Println("Function CancelTransaction")
	w.Header().Set("Content-Type", "application/json")
	connStr := "user=postgres password=alemip10 dbname=BankSystem sslmode=disable"
	var db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println("There was an error connecting to the database")
		defer db.Close()
		err = json.NewEncoder(w).Encode("500 Internal Server Error")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}else{
		log.Println("Database connection established")
	}
	defer db.Close()
	params := mux.Vars(r)
	transaction := settings.Transaction{}
	account := settings.Account{}
	row := db.QueryRow("select * from transactions where id = $1", params["id"])
	err = row.Scan(&transaction.Id, &transaction.Id_acc_reciver, &transaction.Id_acc_sender,
		&transaction.Type, &transaction.Date, &transaction.Amount)
	if err != nil {
		err = json.NewEncoder(w).Encode("No such transaction")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}
	row = db.QueryRow("select * from accounts where id = $1", transaction.Id_acc_sender)
	err = row.Scan(&account.Id, &account.Id_user, &account.Amount)
	if err != nil {
		err = json.NewEncoder(w).Encode("No such transaction")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}
	if transaction.Type == "transfer" {
		if account.Amount > transaction.Amount {
			tx, err := db.Begin()
			if err != nil {
				log.Println("Could not create transactions")
				err = json.NewEncoder(w).Encode("500 Internal Server Error")
				if err != nil {
					log.Panic("Could not return responce")
				}
			}
			{
				stmt, err := tx.Prepare("update accounts set amount=amount + $1 where id = $2;")
				if err != nil {
					tx.Rollback()
					log.Println("Could not execute query to DB")
					defer stmt.Close()
					err = json.NewEncoder(w).Encode("500 Internal Server Error")
					if err != nil {
						log.Panic("Could not return responce")
					}
					return
				}
				defer stmt.Close()
				_, err = stmt.Exec(transaction.Amount, transaction.Id_acc_sender)
				if err != nil {
					tx.Rollback()
					log.Println("Could not execute query to DB")
					defer stmt.Close()
					err = json.NewEncoder(w).Encode("500 Internal Server Error")
					if err != nil {
						log.Panic("Could not return responce")
					}
					return
				}
			}
			{
				stmt, err := tx.Prepare("update accounts set amount=amount - $1 where id = $2;")
				if err != nil {
					tx.Rollback()
					log.Println("Could not execute query to DB")
					defer stmt.Close()
					err = json.NewEncoder(w).Encode("500 Internal Server Error")
					if err != nil {
						log.Panic("Could not return responce")
					}
					return
				}
				defer stmt.Close()
				_, err = stmt.Exec(transaction.Amount, transaction.Id_acc_reciver)
				if err != nil {
					tx.Rollback()
					log.Println("Could not execute query to DB")
					defer stmt.Close()
					err = json.NewEncoder(w).Encode("500 Internal Server Error")
					if err != nil {
						log.Panic("Could not return responce")
					}
					return
				}
			}
			tx.Commit()
			err = json.NewEncoder(w).Encode("Транзакция была успешна отмененна")
			if err != nil {
				log.Panic("Could not return responce")
			}
		} else {
			err = json.NewEncoder(w).Encode("Is not possible to return the money")
			if err != nil {
				log.Panic("Could not return responce")
			}
			return
		}
	} else  {
		err = json.NewEncoder(w).Encode("Is not possible to return the money")
		if err != nil {
			log.Panic("Could not return responce")
		}
		return
	}
	log.Println("The function was successful")
}
