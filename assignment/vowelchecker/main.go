package main

import (
	"fmt"
	"strings"
)

func main() {
    result := vowelChecker()
    fmt.Println(result)
}

func vowelChecker() string {
    fmt.Println(`CHECK YOUR CHARACTER
    Enter a character :
    `)
    var userInput string
    fmt.Scan(&userInput)
	lowercase := strings.ToLower(userInput)

    switch lowercase[0]{
    case 'a', 'e', 'i', 'o', 'u':
        return "it is a vowel"
    default:
        return "it is a consonant"
    }
}
