package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func wordFreq(t string) map[string]int {
	var clean string
	t = strings.ToLower(t)

	for _, r := range t {

		if !unicode.IsPunct(r) {
			clean += string(r)
		}
	}

	words := strings.Fields(t)
	freq := make(map[string]int)

	for _, s := range words {
		freq[s]++
	}
	return freq
}

func isPalindrome(s string) string {
	var clean string

	s = strings.ToLower(s)
	for _, r := range s {

		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			clean += string(r)
		}
	}

	for i, j := 0, len(clean)-1; i < j; i, j = i+1, j-1 {
		if clean[i] != clean[j] {
			return "This is not a palindrome"
		}

	}
	return "This is a palindrome"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter string to count words here:")
	text, _ := reader.ReadString('\n')

	fmt.Println(wordFreq(text))

	fmt.Print("Enter string to Check palindrome:")
	s, _ := reader.ReadString('\n')
	fmt.Print(isPalindrome(s))
}
