package main

import (
	"rutube_test/database"
)

func main() {
	//var userList database.UserList

	//userList.GenerateUserList(3)
	//userList.ChangeUserData(0, database.UserDataForBirth{"Костя", "Черпачок", time.Now()})
	// timeHelper.CheckBirthdays(userList.UserList, time.Now())
	// userList.PrintAllUsers()

	var listManager database.ListManager

	listManager.GenerateUserLists(100, 10)

	listManager.PrintLists()
}
