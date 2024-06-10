package database

import (
	"fmt"
	"sync"
	"time"
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

func (listManager *ListManager) CheckBirthDayForAllLists(currentDay time.Time) {
	var wg sync.WaitGroup
	for _, userList := range listManager.UserLists {
		wg.Add(1)
		CheckBirthdays(userList, currentDay, &wg)
	}
	wg.Wait()
}

func CheckBirthdays(userList UserList, checkDate time.Time, wg *sync.WaitGroup) { //not optimized
	defer wg.Done()
	for _, user := range userList.GetData() {
		year := checkDate.Year()

		//If it's not a leap year and the user's birthday is on the 29th, it will count as the 28th
		if user.BirthDate.Month() == 2 && checkDate.Day() == 29 && !(year%4 == 0 && (year%100 != 0 || year%400 == 0)) {
			if checkDate.Day() == 28 && user.BirthDate.Month() == checkDate.Month() {
				fmt.Printf("С днем рождения, %s!\n", user.Name)
				continue
			}
		}
		if user.BirthDate.Day() == checkDate.Day() && user.BirthDate.Month() == checkDate.Month() {
			fmt.Printf("С днем рождения, %s!\n", user.Name)
		}
	}
}
