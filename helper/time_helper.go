package helper

import (
	"fmt"
	"rutube_test/database"
)

func Time(user database.UserDataForBirth) {
	fmt.Printf("%d - day, %d - mouth\n", user.BirthDate.Day(), user.BirthDate.Month())

}
