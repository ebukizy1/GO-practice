package main

import "fmt"

func main()  {
	ticketPrice()
	
}
func ticketPrice()  {
	fmt.Println(`TICKET PRICE
	Enter your age: `)
	var userInputAge int
	fmt.Scan(& userInputAge)

	var ticketPrice int
	if(userInputAge >= 0 && userInputAge < 12){
		ticketPrice = 5
		fmt.Println("The price of the Ticket is ",ticketPrice,"naira for Children")
	}else if(userInputAge >= 12 && userInputAge <= 64){
		ticketPrice = 10
		fmt.Println("The price of the Ticket is ",ticketPrice,"naira for Adult")
	}else if(userInputAge > 65){
		ticketPrice = 6
		fmt.Println("The price of the Ticket is #",ticketPrice,"naira for children")
	}
}