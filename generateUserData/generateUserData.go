package generateuserdata

import (
	"math/rand"
	"time"
)

func GenerateRandomDate(startYear int, endYear int) time.Time {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	year := rng.Intn(endYear-startYear+1) + startYear
	month := rng.Intn(12) + 1
	var day int

	switch month {
	case 1:
		day = rng.Intn(31) + 1
	case 2:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) { //check for leap year
			day = rng.Intn(29) + 1 //leap year
		} else {
			day = rng.Intn(28) + 1 //not a leap year
		}
	case 3:
		day = rng.Intn(31) + 1
	case 4:
		day = rng.Intn(30) + 1
	case 5:
		day = rng.Intn(31) + 1
	case 6:
		day = rng.Intn(30) + 1
	case 7:
		day = rng.Intn(31) + 1
	case 8:
		day = rng.Intn(31) + 1
	case 9:
		day = rng.Intn(30) + 1
	case 10:
		day = rng.Intn(31) + 1
	case 11:
		day = rng.Intn(30) + 1
	case 12:
		day = rng.Intn(31) + 1
	default:
		day = rng.Intn(28) + 1
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func GenerateUserNameAndSurname() (string, string) {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

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

	return names[rng.Intn(len(names)+1)], surnames[rng.Intn(len(surnames))+1]
}
