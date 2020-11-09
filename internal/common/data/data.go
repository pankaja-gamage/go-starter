package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DbConn *sql.DB

type DbCredentials struct {
	Username string
	Password string
}

type DbClient struct {
	Type        string
	Host        string
	Port        int
	Credentials DbCredentials
	DbName      string
	SslMode     string
}

func InitDb() {
	client := initDbClient()

	db, err := sql.Open(client.Type, client.buildConnString())
	if err != nil {
		panic(err)
	}
	log.Print("connected to DB")
	db.SetMaxOpenConns(2)
	DbConn = db
}

func CleanUp() {
	log.Print("Closing DB connection.")
	if err := DbConn.Close(); err != nil {
		log.Panic("failed to close the DB connection. Quitting.")
	}
}

func initDbClient() DbClient {
	return DbClient{
		Type:        "postgres",
		Host:        "localhost",
		Port:        5432,
		Credentials: readCredentials(),
		DbName:      "go_db",
		SslMode:     "disable",
	}
}

func readCredentials() DbCredentials {
	return DbCredentials{
		Username: os.Getenv("DB_USR"),
		Password: os.Getenv("DB_USR_PW"),
	}
}

func (client DbClient) buildConnString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		client.Host,
		client.Port,
		client.Credentials.Username,
		client.Credentials.Password,
		client.DbName,
		client.SslMode,
	)
}
