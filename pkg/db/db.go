package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)


type DB interface {
	ConnectDB(host string, port int, user, password, dbname string) (*Database, error)
}

type Database struct {
	Conn *sql.DB
}

type DBConfig struct {
	Host string 
	Port int 
	User string 
	Password string 
	Dbname string 
	Sslmode string 
	Sslrootcert string 
	Sslkey string 
	Sslsert string
}


func ConnectDB(config *DBConfig) (*Database, error) {
	psqInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
	config.Host, config.Port, config.User, config.Password, config.Dbname, config.Sslmode)
	
	db, err := sql.Open("postgres", psqInfo)
	if err != nil {
		return nil, err
	}
	
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	return &Database{Conn: db}, nil
}

