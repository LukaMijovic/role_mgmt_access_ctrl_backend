package database

import (
	"database/sql"
	"fmt"

	creds "github.com/LukaMijovic/role-mgmt-access-ctrl/credentials"
	"github.com/LukaMijovic/role-mgmt-access-ctrl/util"
	_ "github.com/lib/pq"
)

var dataBase *sql.DB

func ConnectToDatabase() error {
	creds, err := util.ParseDatabaseCredentials[creds.DBCredential]("credentials/credentials.json")

	if err != nil {
		fmt.Println(err)

	}

	dataSourceName := "user=" + creds.User + " password=" + creds.Password + " dbname=" + creds.DBName + " sslmode=" + creds.SSLMode

	dataBase, err = sql.Open(creds.DBType, dataSourceName)
	//fmt.Println(dataBase.Stats())

	if err != nil {
		panic(err)
	}

	dataBase.SetConnMaxIdleTime(5)
	dataBase.SetMaxOpenConns(10)

	return nil
}

func DisconnectDatabase() error {
	return dataBase.Close()
}
