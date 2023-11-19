package main

import "fmt"

func main ()  {
	
	result := leapYear()
	fmt.Println(result)

}

func leapYear() string {	
	fmt.Println("Enter a year to check if it is a leap year")
	var userInputtedYear int
	fmt.Scan(&  userInputtedYear)
	if userInputtedYear % 4 == 0 && userInputtedYear % 200 == 0 && userInputtedYear % 400 == 0 {
		return "it is a leap year"
	}else if userInputtedYear % 4 == 0 {
		return "it is a leap year"
	}
	return "is is not a leap year"
}