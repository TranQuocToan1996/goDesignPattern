package structuralPatterns

import "fmt"

type UserListProxy struct {
	SomeDatabase           *UserList
	StackCache             UserList
	StackCapacity          int
	DidLastSearchUsedCache bool
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackCapacity {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}
func (t *UserList) addUser(newUser User) {
	*t = append(*t, newUser)
}

func (u *UserListProxy) FindUser(id int32) (User, error) {

	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		u.DidLastSearchUsedCache = true
		return user, nil
	}

	user, err = u.SomeDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}

	u.addUserToStack(user)
	u.DidLastSearchUsedCache = false
	return user, nil
}
