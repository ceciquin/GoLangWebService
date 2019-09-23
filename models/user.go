package models

import (
	"errors"
	"fmt"
)

//User is
type User struct {
	ID        int
	FirstName string
	LastName  string
}

/*variable block*/
var (
	users  []*User // an array that will point to users, it's more efficient as we are not sharing memory thorugh the application
	nextID = 1
)

//GetUsers is
func GetUsers() []*User {
	return users
}

//AddUsers is
func AddUsers(u User) (User, error) {

	if u.ID != 0 {
		return User{}, errors.New("New User must not include an ID or it must set to zero")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

//GetUserByID is
func GetUserByID(id int) (User, error) {

	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}

	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

//RemoveUserByID is
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with id '%v' not found", id)
}
