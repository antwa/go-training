package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()

	fmt.Println("hari ini [Switch]: ", getHariIndo(currentTime))
	fmt.Println("hari ini [map]: ", getHariIndoMap(currentTime))

	// loop
	fmt.Println("looping")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// range loop
	fmt.Println("Range Loop")
	number := []int{1, 2, 3, 4, 5}

	for _, num := range number {
		fmt.Println(num)
	}

	// while loop
	fmt.Println("While Loop")
	w := 1

	for w <= 5 {
		fmt.Println(w)
		w++
	}
}

// switch version
func getHariIndo(dt time.Time) string {

	day := dt.Weekday()
	switch day {
	case time.Sunday:
		return "Minggu"
	case time.Monday:
		return "Senin"
	case time.Friday:
		return "Jumat"
	case time.Saturday:
		return "Sabtu"
	default:
		return ""

	}

}

// map version
func getHariIndoMap(dt time.Time) string {
	day := dt.Weekday().String()

	indoDay := map[string]string{
		"Sunday":    "Minggu",
		"Monday":    "Senin",
		"Tuesday":   "Selasa",
		"Wednesday": "Rabu",
		"Thursday":  "Kamis",
		"Friday":    "Jumat",
		"Saturday":  "Sabtu",
	}
	indonesianDayName := indoDay[day]
	return indonesianDayName
}
