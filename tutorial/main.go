package main

import (
	"fmt"
	"time"

	"example.com/tutorial/age"
	"example.com/tutorial/user"
)

func main() {
	age.Run()

	var fname string
	var lname string
	fmt.Scan(&fname, &lname)
	// user.RegisterUser(user.User{fname:fname , lname:lname,createdAt:time.Now()})
	user.RegisterUser(user.User{
		Fname:     fname,
		Lname:     lname,
		CreatedAt: time.Now(),
	})
	time.Sleep(500)
	fmt.Println(user.GetAllUser())
	fmt.Print((user.GetAllUser()[0]).CreatedAt)

}
