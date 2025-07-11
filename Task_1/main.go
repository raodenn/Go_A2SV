package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(r *bufio.Reader, s string) string {
	fmt.Print(s)
	input, _ := r.ReadString('\n')

	return strings.TrimSpace(input)

}

func createGrade() grade {
	reader := bufio.NewReader(os.Stdin)
	name := getInput(reader, "Enter your name:")
	amount, _ := strconv.Atoi(getInput(reader, "Enter amount of subjects:"))

	g := newgrade(name, amount)

	return g

}
func validNum(r *bufio.Reader) float64 {
	gp, err := strconv.ParseFloat(getInput(r, "Subject Grade?"), 64)

	if err != nil || gp > 4 {
		fmt.Println("you didn't enter a valid grade")
		gp = validNum(r)
	}

	return gp

}

func addSubjects(g grade) {
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < g.no_subjects; i++ {
		name := getInput(reader, "Subject name?")
		gp := validNum(reader)
		fmt.Println("")

		fmt.Println("")

		g.addItems(name, gp)
	}

}

func main() {
	mygrade := createGrade()
	addSubjects(mygrade)

	fmt.Printf(mygrade.format())
}
