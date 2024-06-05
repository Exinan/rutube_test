package main

import (
	"rutube_test/database"
)

func addOne(users []database.UserDataForBirth) []database.UserDataForBirth {
	var tmp database.UserDataForBirth
	tmp.GenerateUserData()
	users = append(users, tmp)
	return users
}

func main() {
	var users []database.UserDataForBirth
	for i := 0; i < 10; i++ {
		users = addOne(users)
	}

	if users == nil {
	}
	for _, user := range users {
		database.PrintUser(user)
	}
}
