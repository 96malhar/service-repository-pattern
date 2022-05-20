package main

import (
	"os"
	"srp/service"
	"srp/store"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "calhounio_demo"
)

func main() {
	userStore, err := store.InitUserStore(user, host, dbname, port)
	defer store.DeleteUserStore(user, host, dbname, port)
	if err != nil {
		panic(err)
	}

	//writer, _ := os.OpenFile("output.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	writer := os.Stdout
	us := service.InitUserService(writer, userStore)

	us.InsertUser("Jon", "Smith", "jonsmith@gmail.com", 26)
	us.InsertUser("Jane", "Smith", "janesmith@gmail.com", 26)
	us.InsertUser("Bob", "Allen", "boballen@gmail.com", 40)
	us.DeleteUserByEmail("xyz@gmail.com")
	us.RetrieveUserByEmail("jonsmith@gmail.com")
	us.RetrieveUsersLessThanAge(10)
	us.DeleteUserByEmail("abc@gmail.com")
	us.DeleteUserByEmail("jonsmith@gmail.com")
	us.InsertUser("Jon", "Smith", "jonsmith@gmail.com", 27)
	us.RetrieveUserByEmail("abc@gmail.com")
}
