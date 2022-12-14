package behavioralPatterns

type UserCollection struct {
	users []*User
}

func (u *UserCollection) CreateIterator() Iterator {
	return &UserIterator{
		users: u.users,
	}
}
