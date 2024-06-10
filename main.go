package main

import (
	"rutube_test/database"
	"time"
)

func main() {
	var listManager database.ListManager

	listManager.GenerateUserLists(100, 10)

	listManager.CheckBirthDayForAllLists(time.Now())
}
