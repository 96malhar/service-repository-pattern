package store

import (
	"database/sql"
	"fmt"
)

func InitUserStore(user, host, dbname string, port int) (*UserStore, error) {
	db, err := createUserDB(user, host, dbname, port)
	if err != nil {
		return nil, err
	}

	sqlStatement := `CREATE table IF NOT EXISTS users (id SERIAL PRIMARY KEY,age INT,
					first_name TEXT,
  					last_name TEXT,
  					email TEXT UNIQUE NOT NULL);
  					`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		return nil, err
	}

	return &UserStore{db: db}, nil
}

func DeleteUserStore(user, host, dbname string, port int) error {
	err := deleteUserDB(user, host, dbname, port)
	if err != nil {
		return err
	}

	return nil
}

func createUserDB(user, host, dbname string, port int) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s sslmode=disable", host, port, user)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE DATABASE " + dbname)
	if err != nil {
		return nil, err
	}

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func deleteUserDB(user, host, dbname string, port int) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s sslmode=disable", host, port, user)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	_, err = db.Exec("DROP DATABASE IF EXISTS " + dbname + " WITH (FORCE)")
	if err != nil {
		return err
	}

	return nil
}
