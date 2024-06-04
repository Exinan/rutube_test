package main

import (
	"rutube_test/database"
)

func main() {
	var user database.UserDataForBirth
	database.GenerateUserData(&user)
	database.PrintUser(user)
}
