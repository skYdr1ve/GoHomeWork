# GoHomeWork
### Crud api
* GetUsers - получаем список всех пользователей.
* GetUser  - получить определённого пользователя. 
* CreateUser - создать пользователя.
* UpdateUser - изменить данные о пользователе .
* DeleteUser - удалить определённого пользователя.

Примеры обращения к Crud api

| Команда       | Type    | URL                     | Body  |
| ------------- |:--------|:-----------------------:|:------|
| GetUsers      | GET     | localhost:8000/users    | null  |
| GetUser       | GET     | localhost:8000/users/7| null |
| CreateUser    | POST    | localhost:8000/users    | "firs_name": "some_name",<br>"last_name": "some_last_name",<br>"phone_number": "some_phone",<br>"email": "some_emaill" |
| UpdateUser    | PUT     | localhost:8000/users7| "firs_name": "some_name",<br>"last_name": "some_last_name",<br>"phone_number": "some_phone",<br>"email": "some_emaill" |
| DeleteUser    | DELETE  | localhost:8000/users{id}| null  |

### Transactions api
* CreateAccount - создание кошелька.
* DeleteAccount  - удаление кошелька. 
* GetBalance - получение  баланса на кошельке.
* MoneyTransfer - передать деньги с одного кошелка на другой. 
* DepositMoney - положить деньги на кошелёк.
* WithdrawalMoney - снять деньги с кошелька.
* GetTransactions - получить историю транзакций.
* CancelTransaction - отменить транзакцию.

Примеры обращения к Transactions api

| Команда       | Type    | URL                     | Body  |
| ------------- |:--------|:-----------------------:|:------|
| CreateAccount | POST    | localhost:8000/accounts/7| "amount": "1000"  |
| DeleteAccount | DELETE  | localhost:8000/accounts/7| null |
| GetBalance    | GET     | localhost:8000/accounts/7| null |
| MoneyTransfer | PUT     | localhost:8000/accounts/2| "id": "1",<br>"amount": "1000" |
| DepositMoney  | PUT     | localhost:8000/accountsdepositMoney| "id": "2",<br>"amount": "1000" |
| WithdrawalMoney| PUT    | localhost:8000/accountswithdrawal| "id": "5",<br>"amount": "1000"  |
| GetTransactions| PUT    | localhost:8000/accountstransactions/5| "date_first": "2019-06-02",<br>"date_second": "2019-06-02"  |
| CancelTransaction| DELETE| localhost:8000/transactions/5| null  |

#### Для кооректной работы сервера нужно подключить базу данных - это следует сделать, подключив dump(dumpBankSystem) командой psql -U postgres -f (your path)\GoHomeWork\dumpBankSystem (для Windows)
