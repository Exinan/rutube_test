package database

import "fmt"

type UserList struct {
	UserList []UserDataForBirth
}

func (userList *UserList) GetId(name string, surname string) []int {
	var result []int
	for i, user := range userList.UserList {
		if user.CheckElement(name, surname) {
			result = append(result, i)
		}
	}
	return result
}

func (userList *UserList) ChangeUserData(id int, newUserData UserDataForBirth) {
	if len(userList.UserList) > id && id >= 0 {
		userList.UserList[id].SetData(newUserData)
	} else {
		fmt.Println("ERROR in ChangeUserData")
	}
}

func (userList *UserList) AddUser(user UserDataForBirth) {
	userList.UserList = append(userList.UserList, user)
}

func (userList *UserList) GenerateUserList(size int) {
	for i := 0; i < size; i++ {
		var user UserDataForBirth
		user.GenerateUserData()
		userList.AddUser(user)
	}
}

func (userList *UserList) CreateUserList(size int) { //create empty user list
	for i := 0; i < size; i++ {
		var user UserDataForBirth
		userList.AddUser(user)
	}
}

func remove(userList []UserDataForBirth, id int) []UserDataForBirth {
	return append(userList[:id], userList[id+1:]...)
}

func (userList *UserList) DeleteUser(id int) {
	if len(userList.UserList) > id && id >= 0 {
		userList.UserList = remove(userList.UserList, id)
	} else {
		fmt.Println("ERROR in DeleteUser")
	}
}

func (userList *UserList) PrintAllUsers() {
	fmt.Println("----------")
	for _, user := range userList.UserList {
		user.PrintUser()
	}
}
