package database

import (
	"fmt"
	generateHelper "rutube_test/generate_helper"
	"time"
)

type UserDataForBirth struct {
	Name      string
	Surname   string
	BirthDate time.Time
}

func (user *UserDataForBirth) setName(name string) {
	user.Name = name
}

func (user *UserDataForBirth) setSurname(surename string) {
	user.Surname = surename
}

func (user *UserDataForBirth) setBirthDate(birthDate time.Time) {
	user.BirthDate = birthDate
}

func SetAllData(user *UserDataForBirth, name string, surname string, birthDate time.Time) {
	user.setName(name)
	user.setSurname(surname)
	user.setBirthDate(birthDate)
}

func GenerateUserData(user *UserDataForBirth) {

	name, surname := generateHelper.GenerateUserNameAndSurname()
	birthDate := generateHelper.GenerateRandomDate(1990, 2006)

	SetAllData(user, name, surname, birthDate)
}

func (user *UserDataForBirth) GenerateUserData() {
	GenerateUserData(user)
}

func PrintUser(user UserDataForBirth) {
	fmt.Printf("Name: %v, Surname: %v, Day: %v, Month: %v\n", user.Name, user.Surname, user.BirthDate.Day(), user.BirthDate.Month())
}
