package main

import (
	"bufio"
	"os"
	"fmt"
)


func main()  {
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of the pizza :")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating : ", input)
}