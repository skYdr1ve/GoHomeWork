package crud

import (
	"GoTask/settings"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const connStr = "user=postgres password=alemip10 dbname=BankSystem sslmode=disable"

func GetUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Function GetUsers")
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
	rows, err := db.Query("select * from users")

	users := []settings.User{}
	for rows.Next() {
		p := settings.User{}
		err = rows.Scan(&p.Id, &p.First_name, &p.Last_name, &p.Phone_number, &p.Email)
		if err != nil {
			log.Println(err)
			continue
		}
		users = append(users, p)
	}

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Panic("Could not return responce")
	}
	log.Println("The function was successful")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Function GetUser")
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
	row := db.QueryRow("select * from users where id = $1", params["id"])
	user := settings.User{}
	err = row.Scan(&user.Id, &user.First_name, &user.Last_name, &user.Phone_number, &user.Email)
	if err != nil {
		log.Println("No such user")
		err = json.NewEncoder(w).Encode(err)
		if err != nil {
			log.Println("Could not return responce")
		}
		return
	} else {
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			log.Panic("Could not return responce")
		}
		log.Println("The function was successful")
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Function CreateUser")
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
	var user settings.User
	err = json.NewDecoder(r.Body).Decode(&user)
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
	_, err = db.Exec("insert into users (first_name, last_name, phone_number, email) "+
		"values ($1, $2, $3, $4)",
		user.First_name, user.Last_name, user.Phone_number, user.Email)
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
	err=json.NewEncoder(w).Encode("User has been created")
	if err != nil {
		log.Panic("Could not return responce")
	}
	log.Println("The function was successful")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Function UpdateUser")
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
	user := settings.User{}
	err = json.NewDecoder(r.Body).Decode(&user)
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
	user.Id = params["id"]
	_,err = db.Exec("update users set first_name=$1, last_name=$2, phone_number=$3, email=$4 where id = $5", user.First_name,
		user.Last_name, user.Phone_number, user.Email, user.Id)
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
	err = json.NewEncoder(w).Encode("User data has been updated")
	if err != nil {
		log.Panic("Could not return responce")
	}
	log.Println("The function was successful")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Function DeleteUser")
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
			defer stmt.Close()
			err = json.NewEncoder(w).Encode("500 Internal Server Error")
			if err != nil {
				log.Panic("Could not return responce")
			}
			return
		}
	}
	{
		stmt, err := tx.Prepare("delete from accounts where id_user = $1")
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
			defer stmt.Close()
			err = json.NewEncoder(w).Encode("500 Internal Server Error")
			if err != nil {
				log.Panic("Could not return responce")
			}
			return
		}
	}
	{
		stmt, err := tx.Prepare("delete from users where id = $1")
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
			defer stmt.Close()
			err = json.NewEncoder(w).Encode("500 Internal Server Error")
			if err != nil {
				log.Panic("Could not return responce")
			}
			return
		}
	}
	tx.Commit()
	log.Println("The function was successful")
}