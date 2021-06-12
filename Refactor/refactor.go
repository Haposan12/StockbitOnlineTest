package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(newFindFirstStringInBracket("haposan (asasa) (aaaaa)"))
}

func oldFindFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str,"(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket,")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1:indexClosingBracketFound-1])
			}else{
				return ""
			}
		}else{
			return ""
		}
	}else{
		return ""
	}
	return ""
}

func newFindFirstStringInBracket(str string) string {
	//variable
	indexFirstBracketFound := strings.Index(str,"(")

	if len(str) > 0 && indexFirstBracketFound >= 0{
		runes := []rune(str)
		wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
		indexClosingBracketFound := strings.Index(wordsAfterFirstBracket,")")

		if indexClosingBracketFound >= 0 {
			runes := []rune(wordsAfterFirstBracket)
			return string(runes[1:indexClosingBracketFound-1])
		}
	}

	return ""
}
