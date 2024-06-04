package main

import (
	"rutube_test/database"
	"rutube_test/helper"
	"time"
)

func main() {
	a := database.UserDataForBirth{Name: "123", Surname: "456", BirthDate: time.Now()}

	helper.Print()
	helper.Time(a)
	helper.Time(a)
}
