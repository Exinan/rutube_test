package main

import (
	"rutube_test/database"
)

func main() {
	users := database.MakeListOfUser(10)
	database.GenerateListOfUser(&users)
	database.PrintList(users)
}
