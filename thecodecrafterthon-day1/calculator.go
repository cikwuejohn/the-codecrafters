package main

import (
	"fmt"
	"strconv"
)

func add(x, y int64) int64 {
	return x + y
}

func mul(x, y int64) int64 {
	return x * y
}

func sub(x, y int64) int64 {
	return x - y
}

func div(x, y int64) int64 {
	return x / y
}

var operator string
var input1 string
var input2 string

func main() {
	/*
		    scanner := bufio.NewScanner(os.Stdin)

		    fmt.Print("Enter input: ")
		    scanner.Scan()
		func getNum1(){
			fmt.Print("Enter first number: ")
			fmt.Scanln(&num1)
		}
		    text := scanner.Text()
		    words := strings.Fields(text)
		    fmt.Println(words)

		    for i := 0; i < len(word); i++{
		        if words[i] == "add"{
		            wor
		        }
		    }
	*/
	for {
		fmt.Print("Enter operator: ")
		fmt.Scanln(&operator)
		fmt.Println()

		if operator == "quit" {
			fmt.Println("You ended the program!")
			break
		}

		if operator == "help" {
			fmt.Print("Valid operation commands include: + or add, - or sub, * or mul, / or div.\n")
			continue
		} else if operator != "add" && operator != "sub" && operator != "mul" && operator != "div" {
			fmt.Println("Invalid input! Type help for command option")
		}

		fmt.Print("Enter first number: ")
		fmt.Scanln(&input1)
		num1, err := strconv.ParseInt(input1, 10, 64)
		if err != nil {
			fmt.Print("Wrong input type: ", err)
			return
		}

		fmt.Print("Enter second number: ")
		fmt.Scanln(&input2)
		num2, err := strconv.ParseInt(input2, 10, 64)
		if err != nil {
			fmt.Print("Wrong input type: ", err)
			return
		}

		switch operator {
		case "+", "add":
			fmt.Print("Result: ", add(num1, num2), "\n")
			fmt.Println()
			continue

		case "-", "sub":
			fmt.Print("Result: ", sub(num1, num2), "\n")
			fmt.Println()
			continue

		case "*", "mul":
			fmt.Print("Result: ", mul(num1, num2), "\n")
			continue
		case "/", "div":
			if num2 == 0 {
				fmt.Println()
				fmt.Println("Don't try that again!")
				fmt.Println()

			} else {
				fmt.Print("Result: ", div(num1, num2), "\n")
				fmt.Println()
				continue
			}

		}
	}
}
