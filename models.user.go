package main

import "errors"

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Mail string `json:"mail"`
}

// For this demo, we're storing the user list in memory
// We also have some users predefined.
// In a real user, this list will most likely be fetched
// from a database. Moreover, in production settings, you should
// store passwords securely by salting and hashing them instead
// of using them as we're doing in this demo
var userList = []user{
	user{ID: 1, Name: "user1", Mail: "user1@test.com"},
}

// Check if the name and password combination is valid
func isUserValid(name string) bool {
	for _, u := range userList {
		if u.Name == name {
			return true
		}
	}
	return false
}

// Check if the supplied name is available
func isNameAvailable(name string) bool {
	for _, u := range userList {
		if u.Name == name {
			return false
		}
	}
	return true
}

// Return a list of all the users
func getAllUsers() []user {
	return userList
}

// Fetch an user based on the ID supplied
func getUserByID(id int) (*user, error) {
	for _, a := range userList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("User not found")
}

// Create a new user with the title and content provided
func createNewUser(name, mail string) (*user, error) {
	// Set the ID of a new user to one more than the number of users
	a := user{ID: len(userList) + 1, Name: name, Mail: mail}

	// Add the user to the list of users
	userList = append(userList, a)

	return &a, nil
}
