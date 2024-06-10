package generateHelper

import (
	"math/rand"
	"time"
)

func GenerateRandomDate(startYear int, endYear int) time.Time {

	year := rand.Intn(endYear-startYear+1) + startYear
	month := rand.Intn(12) + 1
	var day int

	switch month {
	case 1:
		day = rand.Intn(31) + 1
	case 2:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) { //check for leap year
			day = rand.Intn(29) + 1 //leap year
		} else {
			day = rand.Intn(28) + 1 //not a leap year
		}
	case 3:
		day = rand.Intn(31) + 1
	case 4:
		day = rand.Intn(30) + 1
	case 5:
		day = rand.Intn(31) + 1
	case 6:
		day = rand.Intn(30) + 1
	case 7:
		day = rand.Intn(31) + 1
	case 8:
		day = rand.Intn(31) + 1
	case 9:
		day = rand.Intn(30) + 1
	case 10:
		day = rand.Intn(31) + 1
	case 11:
		day = rand.Intn(30) + 1
	case 12:
		day = rand.Intn(31) + 1
	default:
		day = rand.Intn(28) + 1
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func GenerateUserNameAndSurname() (string, string) {

	names := []string{"Иван", "Елена", "Александр", "Мария", "Дмитрий", "Анна",
		"Сергей", "Ольга", "Андрей", "Татьяна", "Михаил", "Екатерина", "Алексей",
		"Наталья", "Константин", "Ирина", "Юрий", "Виктория", "Павел", "Светлана",
		"Владимир", "Людмила", "Артем", "Юлия", "Игорь", "Евгения", "Максим",
		"Анастасия", "Николай", "Валентина"}
	surnames := []string{"Смирнов", "Иванов", "Кузнецов", "Соколов", "Попов",
		"Лебедев", "Козлов", "Новиков", "Морозов", "Петров", "Волков", "Соловьев",
		"Васильев", "Зайцев", "Павлов", "Семенов", "Голубев", "Виноградов", "Богданов",
		"Воробьев", "Федоров", "Михайлов", "Беляев", "Тарасов", "Белов", "Комаров",
		"Орлов", "Киселев", "Макаров", "Андреев"}

	return names[rand.Intn(len(names))], surnames[rand.Intn(len(surnames))]
}
