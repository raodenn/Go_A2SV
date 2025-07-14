package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

func StartConsole(lib *services.Library) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Library Menu ---")
		fmt.Println("1. Add Book")
		fmt.Println("2. Add Member")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("0. Exit")
		fmt.Print("Select an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			handleAddBook(lib, reader)
		case "2":
			handleAddMember(lib, reader)
		case "3":
			handleBorrowBook(lib, reader)
		case "4":
			handleReturnBook(lib, reader)
		case "5":
			handleListAvailableBooks(lib)
		case "6":
			handleListBorrowedBooks(lib, reader)
		case "0":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}

// prints all current members to show user available id's
func printMembers(members []*models.Member) {
	if len(members) == 0 {
		fmt.Println("No members found.")
		return
	}
	for _, m := range members {
		fmt.Printf("ID: %d | Name: %s\n", m.Id, m.Name)
	}
}

// check if input is a valid input
func checkIntInput(reader *bufio.Reader, prompt string) (int, error) {
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)

	id, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %v", err)
	}
	return id, nil
}

// handles input for create book
func handleAddBook(lib *services.Library, reader *bufio.Reader) {
	fmt.Print("Enter book title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter author name: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	lib.CreateBook(author, title)
	fmt.Println("Book added successfully.")
}

// handles input for adding book

func handleAddMember(lib *services.Library, reader *bufio.Reader) {
	fmt.Print("Enter member name:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	lib.CreateMember(name)
	fmt.Println("Member add successfully.")
}

// handles input for borrowing book

func handleBorrowBook(lib *services.Library, reader *bufio.Reader) {
	fmt.Println("Available Books:")
	availableBooks := lib.ListAvailableBooks()
	printBooks(availableBooks)

	fmt.Println("Registered Members:")
	members := lib.ListMembers()
	printMembers(members)

	bookID, err := checkIntInput(reader, "Enter Book ID: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	memberID, err := checkIntInput(reader, "Enter Member ID: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = lib.BorrowBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book borrowed successfully.")
	}
}

// handles input for returning book

func handleReturnBook(lib *services.Library, reader *bufio.Reader) {
	fmt.Println("Borrowed Books:")
	fmt.Println("Registered Members:")
	members := lib.ListMembers()
	printMembers(members)

	memberID, err := checkIntInput(reader, "Enter member ID: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	borrowedBooks, _ := lib.ListBorrowedBooks(memberID)
	printBooks(borrowedBooks)

	bookID, err := checkIntInput(reader, "Enter book ID: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = lib.ReturnBook(bookID, memberID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Book returned successfully.")
	}
}

// handles input for printing books to console

func printBooks(books []*models.Book) {
	if len(books) == 0 {
		fmt.Println("No books found.")
	} else {
		for _, book := range books {
			fmt.Printf(" ID:%d | Title: %s | Author:%s | Status:%s\n",
				book.Id, book.Title, book.Author, book.Status)
		}
	}
}

func handleListAvailableBooks(lib *services.Library) {
	books := lib.ListAvailableBooks()
	printBooks(books)

}

// handles input for listing borrowed books

func handleListBorrowedBooks(lib *services.Library, reader *bufio.Reader) {
	fmt.Println("Registered Members:")
	members := lib.ListMembers()
	printMembers(members)

	memberId, err := checkIntInput(reader, "Enter member Id:")
	if err != nil {
		fmt.Println(err)
		return
	}
	books, err := lib.ListBorrowedBooks(memberId)
	if err != nil {
		fmt.Println(err)
	} else {
		printBooks(books)
	}

}
