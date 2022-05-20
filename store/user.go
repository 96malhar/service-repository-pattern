package store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"srp/models"
)

type Database interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type UserStore struct {
	db Database
}

func (s *UserStore) InsertUser(firstName, lastName, email string, age int) (int, error) {
	sqlStatement := `INSERT INTO users (age, email, first_name, last_name)
					 VALUES ($1, $2, $3, $4)
					 RETURNING id;`

	id := 0
	err := s.db.QueryRow(sqlStatement, age, email, firstName, lastName).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *UserStore) DeleteUserByEmail(email string) (int64, error) {
	sqlStatement := `DELETE FROM users WHERE email = $1;`
	result, err := s.db.Exec(sqlStatement, email)
	if err != nil {
		return 0, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (s *UserStore) RetrieveUserByEmail(email string) (*models.User, error) {
	sqlStatement := `SELECT * FROM users WHERE email = $1;`
	var user models.User

	err := s.db.QueryRow(sqlStatement, email).Scan(&user.Id, &user.Age, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserStore) RetrieveUsersLessThanAge(age int) ([]*models.User, error) {
	sqlStatement := `SELECT * FROM users WHERE age <= $1`
	rows, err := s.db.Query(sqlStatement, age)
	if err != nil {
		return nil, err
	}

	var users []*models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Age, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
