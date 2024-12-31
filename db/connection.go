package db

import (
	"database/sql"
	"log"

	"github.com/BraianMarinP/todo-backend/config"
	_ "github.com/go-sql-driver/mysql"
)

var databaseConnection = ConnectDB()

func ConnectDB() *sql.DB {

	config.LoadConfig()

	dbHost := config.GetEnviromentVariable("DB_HOST")
	dbPort := config.GetEnviromentVariable("DB_PORT")
	dbUser := config.GetEnviromentVariable("DB_USER")
	dbPassword := config.GetEnviromentVariable("DB_PASSWORD")
	dbSchema := config.GetEnviromentVariable("DB_SCHEMA")

	// Perfom a connection to the database 'USER:PASSWORD@tcp(HOST:PORT)/SCHEMA'
	connectionString := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbSchema
	connection, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err.Error())
		return connection
	}

	// If an error has ocurred, it returns an error
	err = connection.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("The database connection was successful.")
	return connection
}
