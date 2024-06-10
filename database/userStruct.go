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

func (user *UserDataForBirth) SetAllData(name string, surname string, birthDate time.Time) {
	user.setName(name)
	user.setSurname(surname)
	user.setBirthDate(birthDate)
}

func (user *UserDataForBirth) SetData(newUserData UserDataForBirth) {
	user.SetAllData(newUserData.Name, newUserData.Surname, newUserData.BirthDate)
}

func (user *UserDataForBirth) GenerateUserData() {
	name, surname := generateHelper.GenerateUserNameAndSurname()
	birthDate := generateHelper.GenerateRandomDate(1990, 2006)

	user.SetAllData(name, surname, birthDate)
}

func (user *UserDataForBirth) PrintUser() {
	fmt.Printf("Name: %v, Surname: %v, Day: %v, Month: %v\n", user.Name, user.Surname, user.BirthDate.Day(), user.BirthDate.Month())
}

func (user *UserDataForBirth) CheckElement(name string, surname string) bool {
	return user.Name == name && user.Surname == surname
}

func (user *UserDataForBirth) GetBithDate() time.Time {
	return user.BirthDate
}
