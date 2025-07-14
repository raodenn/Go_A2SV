# Library Management System Documentation

## Overview
The Library Management System is a console-based Go application designed to manage books and members. It allows adding, borrowing, returning, and listing books, as well as managing members. The system uses Go’s core features such as structs, interfaces, slices, maps, and packages to demonstrate modular design and basic object-oriented principles in Go.

---

## Project Structure

library_management/
├── main.go                        // Application entry point
├── controllers/                  
│   └── library_controller.go     // Console interaction logic
├── models/                      
│   ├── book.go                   // Book struct and dummy book data
│   └── member.go                 // Member struct and dummy member data
├── services/                    
│   └── library_service.go        // Business logic and interface implementation
├── docs/
│   └── documentation.md          // You are here
└── go.mod                        // Go module definition

---

## Packages

### models  
Defines core data structures used in the application:

```go
type Book struct {
    Id     int
    Title  string
    Author string
    Status string // "Available" or "Borrowed"
}

type Member struct {
    Id            int
    Name          string
    BorrowedBooks []*Book
}
```
Also includes DummyBooks and DummyMembers for pre-populated data.

## services

### LibraryManager Interface

Defines the abstraction layer for the librarys operations:
```go
type LibraryManager interface {
    AddBook(book *models.Book)
    RemoveBook(bookID int)
    BorrowBook(bookID int, memberID int) error
    ReturnBook(bookID int, memberID int) error
    ListAvailableBooks() []*models.Book
    ListBorrowedBooks(memberID int) ([]*models.Book, error)
    ListMembers() []*models.Member
}
```
## Library Struct
Implements the LibraryManager interface. Internally stores:
```go
type Library struct {
    Books        map[int]*models.Book
    Members      map[int]*models.Member
    NextBookID   int
    NextMemberID int
}
```
Key methods include:

-CreateBook(author, title string)

-CreateMember(name string)

-BookExistsAndAvailable(id int) (*Book, error)

-MemberExists(id int) (*Member, error)

And all interface methods listed above.

## controllers
Handles user interactions via the console.

Key responsibilities:

-Displaying menus

-Collecting user input

-Validating inputs

-Invoking service methods

### Important functions:

StartConsole(lib LibraryManager)

handleAddBook()

handleBorrowBook()

handleReturnBook()

handleListAvailableBooks()

