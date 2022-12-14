package behavioralPatterns

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	user1 := &User{
		Name: "a",
		Age:  30,
	}
	user2 := &User{
		Name: "b",
		Age:  20,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}
}
