package structuralPatterns

import "fmt"

// UserFinder is the interface that the database and the Proxy implement
type UserFinder interface {
	FindUser(id int32) (User, error)
}

type User struct {
	ID int32
}

type UserList []User

// Proxy uses the same interface as the type it wraps
func (u *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*u); i++ {
		if (*u)[i].ID == id {
			return (*u)[i], nil
		}
	}
	return User{}, fmt.Errorf("user %v could not be found", id)
}
