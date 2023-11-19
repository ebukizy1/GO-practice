package main

import (

	"fmt"
)

func main()  {
	loopActive := true

displayPizzaOption(&loopActive)
	
	

}


func displayPizzaOption(loopActive *bool) {
	
	for  *loopActive {
		fmt.Println(`Welcome to Emax pizza store:
			 press 1 --> view menu 
			 press 0 --> To exit
		`)
		var userInput int
		fmt.Scan(&userInput)
		switch userInput {
		case 1:
			viewMenu(loopActive)
		case 0:
			exist(loopActive)
		default:
			fmt.Println("Invalid input. Please try again.")
		}
	}
}

func viewMenu(loopActive *bool) {
	fmt.Println(`Emax Pizza Menu
		press 1 --> chicken pie = 5000
		press 2 --> potato pizza pie = 10000
		press 3 --> turkey pie = 8500
		press 4 --> chicken pie with drink = 6900
		press 6 --> potato pie with drink and chicken = 1500
		press 0 --> To exit 
	`)
	var userInput int
	fmt.Scan(&userInput)
	switch userInput {
	case 1, 2, 3, 4, 6:
		OrderResult(loopActive)
	case 0:
		exist(loopActive)
	default:
		fmt.Println("Invalid input. Please try again.")
	}
}

func OrderResult(loopActive *bool) {
	fmt.Println(`
		=========================

				ORDER SUCCESSFUL
				   will be 
				  available
				 in 10 minutes

		========================
	`)
	fmt.Println(`Do you want to continue
		press 1 --> view menu 
		press 0 --> To exit
	`)
	var userInput int
	fmt.Scan(&userInput)
	switch userInput {
	case 1:
		viewMenu(loopActive)
	case 0:
		exist(loopActive)
	default:
		fmt.Println("Invalid input. Please try again.")
	}
}

func exist(loopActive *bool) {
	fmt.Println("THANKS FOR YOUR PATRONAGE")
	*loopActive = false
}
