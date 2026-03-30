// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [Your Name]
// Squad:  [Your Squad Name]
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func toUppercase(str string) string {
	return strings.ToUpper(str)
}
func toLowercase(str string) string {
	return strings.ToLower(str)
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

// func reverse(str string) string{
// 	words := strings.Fields(str)
// 	for i := 0; i < len(words); i++{
// 		words[i]
// 	}
// }

func main() {
	fmt.Println(toUppercase("sentinel is online"))
	fmt.Println(doTitle("the fall of the western power grid"))
	fmt.Println(toSnake("Operation Gopher Protocol"))
	fmt.Println(doTitle("a threat in the north"))

}
