package database

import (
	"fmt"
	//"rutube_test/generateuserdata"
	"time"
)

type UserDataForBirth struct {
	Name      string
	Surname   string
	BirthDate time.Time
}

func MakeListOfUser(count int) []UserDataForBirth {
	return make([]UserDataForBirth, count)
}

func GenerateUserData(user *UserDataForBirth) {
	//user.BirthDate = generateuserdata.GenerateRandomDate()
	//random initialization of names and surnames from the lists above
	// name, surname := generateuserdata.GenerateUserNameAndSurname()
	// user.Name = name
	// user.Surname = surname

	//random initialization of date
	//user.BirthDate = generateuserdata.GenerateRandomDate(1990, 2006)
}

func GenerateListOfUser(users *[]UserDataForBirth) {
	for _, user := range *users {
		GenerateUserData(&user)
	}
}

func PrintList(users []UserDataForBirth) {
	for _, user := range users {
		fmt.Println(user)
		//fmt.Printf("Name: %d, Surname: %d, Day: %d, Month: %d", user.Name, user.Surname, user.BirthDate.Day(), user.BirthDate.Month())
	}
}
