package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatal("This program expects 3 arguments - an operator and two numbers")
	}

	var operator string
	operator = os.Args[1]

	var num1, num2 int64
	_, err := fmt.Sscanf(os.Args[2], "%d", &num1)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Sscanf(os.Args[3], "%d", &num2)
	if err != nil {
		log.Fatal(err)
	}

	if operator == "+" {
		fmt.Println(num1 + num2)
	} else if operator == "*" {
		fmt.Println(num1 * num2)
	} else {
		log.Fatal("Unknown operator: " + operator)
	}
}
