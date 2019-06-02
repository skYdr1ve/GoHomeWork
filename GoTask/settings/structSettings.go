package settings

type User struct {
	Id           string `json:"id"`
	First_name   string `json:"firs_name"`
	Last_name    string `json:"last_name"`
	Phone_number string `json:"phone_number"`
	Email        string `json:"email"`
}

type Account struct {
	Id      string `json:"id"`
	Id_user string `json:"id_user"`
	Amount  string `json:"amount"`
}

type Transaction struct {
	Id             string `json:"id"`
	Id_acc_reciver string `json:"id_acc_reciver"`
	Id_acc_sender  string `json:"id_acc_sender"`
	Type           string `json:"type"`
	Date           string `json:"date"`
	Amount         string `json:"amount"`
}

type Date struct {
	Date_first  string `json:"date_first"`
	Date_second string `json:"date_second"`
}