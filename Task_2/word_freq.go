package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func removePunc(s string) string {
	var res string

	for _, r := range s {
		if !unicode.IsPunct(r) {
			res += string(r)
		}
	}
	return res
}

func wordFreq(t string) map[string]int {
	t = strings.ToLower(t)
	t = removePunc(t)
	words := strings.Fields(t)
	freq := make(map[string]int)

	for _, s := range words {
		freq[s]++
	}
	return freq
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter string here:")
	text, _ := reader.ReadString('\n')

	fmt.Print(wordFreq(text))

}
