package database

import (
	"fmt"
	"sync"
)

type ListManager struct {
	UserLists []UserList
}

func (listManager *ListManager) AddUserList(userList UserList) {
	listManager.UserLists = append(listManager.UserLists, userList)
}

func removeList(userLists []UserList, id int) []UserList {
	return append(userLists[:id], userLists[id+1:]...)
}

func (listManager *ListManager) DeleteList(id int) {
	if len(listManager.UserLists) > id && id >= 0 {
		listManager.UserLists = removeList(listManager.UserLists, id)
	} else {
		fmt.Println("ERROR in DeleteUser")
	}
}

func (listManager *ListManager) PrintLists() {
	for _, userList := range listManager.UserLists {
		userList.PrintAllUsers()
	}
}

func (listManager *ListManager) GenerateUserLists(n int, size int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go listManager.GenerateUserList(size, &wg)
	}
	wg.Wait()
}

func (listManager *ListManager) GenerateUserList(size int, wg *sync.WaitGroup) {
	defer wg.Done()
	var tmpUserList UserList
	tmpUserList.GenerateUserList(size)
	listManager.AddUserList(tmpUserList)
}
