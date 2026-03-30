package main

import (
	"fmt"
	"strconv"
)

func hextoDec(numStr string) int64 {
	n, err := strconv.ParseInt(numStr, 16, 64)
	if err != nil {
		fmt.Print(numStr, " is not a valid hex! \n")
		return n
	} else {
		fmt.Print("Decimal: ", n, "\n")
		return n
	}
}

func bintoDec(numStr string) int64 {
	n, err := strconv.ParseInt(numStr, 2, 64)
	if err != nil {
		fmt.Print(numStr, " is not a valid binary! \n")
		return n
	} else {
		fmt.Print("Decimal: ", n, "\n")
		return n
	}
}

func dectoHex(numStr string) string {
	n, err := strconv.ParseInt(numStr, 10, 64)
	result := strconv.FormatInt(n, 16)

	if err != nil {
		return result
	} else {
		fmt.Print("Hex: ", result, "\n")
		return result
	}
}

func dectoBin(numStr string) string {
	n, err := strconv.Atoi(numStr)
	result := strconv.FormatInt(int64(n), 2)
	if err != nil {
		fmt.Print(numStr, " is not a valid dec!\n")
		return result
	} else {
		fmt.Print("Bin: ", result, "\n")
		return result
	}
}

func main() {

	for {
		var numStr string
		var baseConv string

		fmt.Print("What conversion do you want to do: ")
		fmt.Scanln(&baseConv)
		if baseConv == "help" {
			fmt.Print("Possible conversions commands: dec, hex & bin. Enter quit to end.\n", "\n")
			continue
		}

		if baseConv == "quit" {
			fmt.Println()
			fmt.Println("You ended the program!")
			return
		}

		fmt.Print("Input base to convert: ")
		fmt.Scanln(&numStr)

		switch baseConv {
		case "hex":
			fmt.Println()
			hextoDec(numStr)
			fmt.Println()

		case "bin":
			fmt.Println()
			bintoDec(numStr)
			fmt.Println()

		case "dec":
			fmt.Println()
			dectoHex(numStr)

			dectoBin(numStr)
			fmt.Println()

		}
	}

}
