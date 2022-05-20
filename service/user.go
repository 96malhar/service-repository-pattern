package service

import (
	"fmt"
	"io"
	"srp/models"
)

type Repository interface {
	InsertUser(firstName, lastName, email string, age int) (int, error)
	DeleteUserByEmail(email string) (int64, error)
	RetrieveUserByEmail(email string) (*models.User, error)
	RetrieveUsersLessThanAge(age int) ([]*models.User, error)
}

type UserService struct {
	repo Repository
	w    io.Writer
}

func InitUserService(w io.Writer, repository Repository) *UserService {
	fmt.Fprintf(w, "Successfully initiated user service with user store\n")
	return &UserService{repo: repository, w: w}
}

func (s *UserService) Print(u models.User) {
	fmt.Fprintf(s.w, "USER: ID = %d, FirstName = %s, LastName = %s, Age = %d, Email = %s\n",
		u.Id, u.FirstName, u.LastName, u.Age, u.Email)
}

func (s *UserService) InsertUser(firstName, lastName, email string, age int) {
	id, err := s.repo.InsertUser(firstName, lastName, email, age)
	if err != nil {
		fmt.Fprintln(s.w, "Failed to insert new user due to the following error:", err)
		return
	}
	fmt.Fprintln(s.w, "Successfully inserted user with ID:", id)
}

func (s *UserService) DeleteUserByEmail(email string) {
	n, err := s.repo.DeleteUserByEmail(email)
	if err != nil {
		panic(err)
	}
	if n == 0 {
		fmt.Fprintln(s.w, "Could not find any user with email:", email)
		return
	}
	fmt.Fprintln(s.w, "Successfully deleted user with email:", email)
}

func (s *UserService) RetrieveUserByEmail(email string) {
	u, err := s.repo.RetrieveUserByEmail(email)
	if err != nil {
		fmt.Fprintln(s.w, "Could not find any user with email:", email)
		return
	}
	s.Print(*u)
}

func (s *UserService) RetrieveUsersLessThanAge(age int) {
	users, err := s.repo.RetrieveUsersLessThanAge(age)
	if err != nil {
		fmt.Fprintln(s.w, "Failed to find user below age:", age)
		return
	}

	if len(users) == 0 {
		fmt.Fprintln(s.w, "No users exist below age", age)
		return
	}

	fmt.Fprintf(s.w, "\nUsers below %d years of age:\n", age)
	for _, u := range users {
		s.Print(*u)
	}
}
