package main

import (
	"rutube_test/database"
	"time"
)

func main() {
	var userList database.UserList

	userList.GenerateUserList(10)
	userList.ChangeUserData(0, database.UserDataForBirth{"Костя", "Черпачок", time.Now()})
	userList.PrintAllUsers()
	userList.DeleteUser(0)
	userList.PrintAllUsers()
}
