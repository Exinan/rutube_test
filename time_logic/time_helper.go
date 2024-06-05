package timeHelper

import (
	"fmt"
	"rutube_test/database"
	"time"
)

func CheckBirthdays(users []database.UserDataForBirth, checkDate time.Time) { //not optimized
	for _, user := range users {
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
