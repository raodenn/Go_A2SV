package services

import (
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book *models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []*models.Book
	ListBorrowedBooks(memberID int) ([]*models.Book, error)
}

type Library struct {
	Books        map[int]*models.Book
	Members      map[int]*models.Member
	NextBookID   int
	NextMemberID int
}

// method to initialize library
func NewLibrary() *Library {
	lib := &Library{
		Books:        make(map[int]*models.Book),
		Members:      make(map[int]*models.Member),
		NextBookID:   1,
		NextMemberID: 1,
	}
	for _, book := range models.DummyBooks {
		b := book
		lib.Books[b.Id] = &b
		if b.Id >= lib.NextBookID {
			lib.NextBookID = b.Id + 1
		}
	}

	for _, member := range models.DummyMembers {
		m := member
		lib.Members[m.Id] = &m
		if m.Id >= lib.NextMemberID {
			lib.NextMemberID = m.Id + 1
		}
	}
	return lib
}

// adds book to library
func (lib *Library) AddBook(book *models.Book) {

	lib.Books[book.Id] = book
}

// removes books from library
func (lib *Library) RemoveBook(bookID int) {
	delete(lib.Books, bookID)
}

// borrow book for a member
// checks if both book and member exist returns error if they don't
func (lib *Library) BorrowBook(bookID int, memberID int) error {
	book, err := lib.BookExistsAndAvailable(bookID)
	if err != nil {
		return err
	}
	member, err := lib.MemberExists(memberID)

	if err != nil {
		return err
	}

	book.Status = "Borrowed"
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	return nil
}

// return book for a member
// checks if book is borrowed and if member exists returns error if they don't
func (lib *Library) ReturnBook(bookID int, memberID int) error {
	book, err := lib.BookExistsAndBorrowed(bookID)

	if err != nil {
		return err
	}

	member, err := lib.MemberExists(memberID)

	if err != nil {
		return err
	}
	book.Status = "Available"
	member.BorrowedBooks = RemoveBookunordered(member.BorrowedBooks, book.Id)
	return nil
}

// returns a slice of all available books
func (lib *Library) ListAvailableBooks() []*models.Book {
	var availableBooks []*models.Book
	for _, book := range lib.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

// returns a slice of all the borrowed books for a member with id
func (lib *Library) ListBorrowedBooks(memberID int) ([]*models.Book, error) {
	var borrowedBooks []*models.Book
	member, err := lib.MemberExists(memberID)
	if err != nil {
		return nil, err
	}

	for _, book := range member.BorrowedBooks {
		if book.Status == "Borrowed" {
			borrowedBooks = append(borrowedBooks, book)
		}
	}
	return borrowedBooks, nil
}
