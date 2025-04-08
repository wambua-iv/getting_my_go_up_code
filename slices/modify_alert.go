package main

// Sample program to show how one needs to be careful when appending
// to a slice when you have a reference to an element.

import "fmt"

type user struct {
	likes int
}

func Modify() {

	// Declare a slice of 3 users.
	// users := make([]user, 3, 5)
	users := make([]user, 3)

	// Share the user at index 1.
	shareUser := &users[1]

	// Add a like for the user that was shared.
	shareUser.likes++

	// Display the number of likes for all users.
	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}

	// Add a new user.
	// if the capacity of the backing array does not allow for adding elements
	// a new array is added and all modifications to existing data maybe lost
	users = append(users, user{})

	// Add another like for the user that was shared.
	shareUser.likes++

	// Display the number of likes for all users.
	fmt.Println("*************************")
	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}

	// Notice the last like has not been recorded.
}