package user

import (
	"fmt"
	"time"
)

type User struct {
	Fname     string
	Lname     string
	CreatedAt time.Time
}

var all_user []User

func RegisterUser(u User) {
	fmt.Println("registering new user")
	all_user = append(all_user, u)

}
func GetAllUser() []User {
	return all_user
}

// makibg above to method
func (u *User) GetAllUserMethod() []User {
	return all_user
}
