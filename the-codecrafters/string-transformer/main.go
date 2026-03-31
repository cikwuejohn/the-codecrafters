// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Christopher John]
// Squad:  [The Compilers]
package main

import (
    "fmt"
    "strings"
    "unicode"
    "os"
    "bufio"
)

func toUppercase(str string) string {
    return strings.Join(strings.Fields(strings.ToUpper(str)), " ")
}
func toLowercase(str string) string {
    return strings.Join(strings.Fields(strings.ToLower(str)), " ")
}
func toCaps(str string) string {
    words := strings.Fields(str)
    for i := 0; i < len(words); i++ {
        words[i] = helpCap(words[i])
    }
    return strings.Join(words, " ")
}
func helpCap(str string) string {
    title := strings.ToUpper(str[:1]) + strings.ToLower(str[1:])
    return title
}

func wordsCheck(str string) bool {
    smallWords := []string{"a", "an", "the", "and", "but", "or", "for", "nor", "on", "at", "to", "by", "in", "of", "up", "as", "is", "it"}

    for i := 0; i < len(smallWords); i++ {
        if strings.Contains(smallWords[i], str) {
            return true
        }
    }
    return false
}
func doTitle(str string) string {
    words := strings.Fields(str)
    for i := 0; i < len(words); i++ {
        if wordsCheck(words[i]) == true && i > 0 {
            words[i] = strings.ToLower(words[i])
        } else {
            words[i] = helpCap(words[i])
        }
    }
    return strings.Join(words, " ")
}

func toSnake(str string) string {
    var cleaned strings.Builder

    for _, char := range str {
        if unicode.IsLetter(char) || unicode.IsDigit(char) || unicode.IsSpace(char) {
            cleaned.WriteRune(char)
        }
    }
    words := strings.ToLower(cleaned.String())

    return strings.Join(strings.Fields(words), "_")
}

func reverseWord(str string) string{
    words := strings.Fields(str)
    for i := 0; i < len(words); i++{
        words[i] = swapChar(words[i])
    }
    return strings.Join(words, " ")
}

func swapChar(str string) string{
    runes := []rune(str)
    l := len(runes)
    
    for i := 0; i < l/2; i++{
        runes[i], runes[l-1-i] = runes[l-1-i], runes[i]
    }
    return string(runes)
}
func checkOps(str string) bool{
	ws := "upper lower caps snake title reverse"
	slice := strings.Fields(ws)

	for i := 0; i < len(slice); i++{
		if slice[i] == str{
			return true
		}
	}
	return false
}
func listOps(){
	ws := "upper lower caps snake title reverse"
	list := strings.Fields(ws)
	for _, w:= range list{
		fmt.Println(w)
	}
}
var operator string

func main() {
    scanner := bufio.NewScanner(os.Stdin)

	for{
    fmt.Print("Enter conversion command: ")
    fmt.Scanln(&operator)

	if operator == "exit"{
		fmt.Print("\nYou ended the program!\n")
		return
	}

	if !checkOps(operator){
		fmt.Print("\n", operator, " is not a valid command!\n\n")
		fmt.Println("Possible commands include:")
		listOps()
		continue
	}

    fmt.Print("Enter text to be transformed: ")
    scanner.Scan()

    text := scanner.Text()

	if len(text) == 0 || text == " "{
		fmt.Print("No text provided! Usage: cap <text>")
	}


    switch operator {
    case "upper":
        fmt.Print("\n", toUppercase(text), "\n")
		continue
        
	case "lower":
		fmt.Print("\n", toLowercase(text), "\n")
		continue

	case "cap":
		fmt.Print("\n", toCaps(text), "\n")
		continue

	case "title":
		fmt.Print("\n", doTitle(text), "\n")
		continue

	case "snake":
		fmt.Print("\n", toSnake(text), "\n")
		continue

	case "reverse":
		fmt.Print("\n", reverseWord(text), "\n")
		continue

    }
    
    

}
}

