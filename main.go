package main

import (
	"fmt"
	//"rutube_test/database"
	"rutube_test/generateuserdata"
)

func main() {
	// users := database.MakeListOfUser(10)
	// database.GenerateListOfUser(&users)
	// database.PrintList(users)

	for i := 0; i < 10; i++ {
		fmt.Println(generateuserdata.GenerateUserNameAndSurname())
	}
}
