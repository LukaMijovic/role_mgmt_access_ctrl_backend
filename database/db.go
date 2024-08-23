package database

import "database/sql"

var dataBase *sql.DB
var creds *DBCredential = 

func ConnectToDatabase() {
	var err error

	database, err = sql.Open()
}
