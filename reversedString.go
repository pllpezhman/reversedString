package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    reader := bufio.NewReader((os.Stdin))
    fmt.Print("Enter a sentence: ")
    sentence, _ := reader.ReadString('\n')
    sentence = strings.TrimSpace(sentence)
    reversedSentence := reverseString(sentence)
    fmt.Println(reversedSentence)
}

func reverseString(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}