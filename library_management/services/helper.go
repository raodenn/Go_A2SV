package services

import (
	"fmt"
	"library_management/models"
)

// helper functions
func (lib *Library) CreateBook(author string, title string) {
	book := models.Book{
		Id:     lib.NextBookID,
		Title:  title,
		Author: author,
		Status: "Available",
	}
	lib.NextBookID++
	lib.AddBook(&book)
}

func (lib *Library) BookExistsAndBorrowed(bookID int) (*models.Book, error) {
	book, exists := lib.Books[bookID]
	if !exists {
		return nil, fmt.Errorf("book with ID %d not found", bookID)
	}
	if book.Status != "Borrowed" {
		return nil, fmt.Errorf("book with ID %d is not currently borrowed", bookID)
	}
	return book, nil
}

func (lib *Library) BookExistsAndAvailable(bookID int) (*models.Book, error) {
	book, exists := lib.Books[bookID]
	if !exists {
		return nil, fmt.Errorf("book with ID %d not found", bookID)
	}
	if book.Status != "Available" {
		return nil, fmt.Errorf("book with ID %d is not available", bookID)
	}
	return book, nil
}

func (lib *Library) MemberExists(memberID int) (*models.Member, error) {
	member, exists := lib.Members[memberID]
	if !exists {
		return nil, fmt.Errorf("member with ID %d doesn't exist", memberID)
	}
	return member, nil
}

func RemoveBookunordered(slice []*models.Book, bookID int) []*models.Book {
	for i, b := range slice {
		if b.Id == bookID {
			slice[i] = slice[len(slice)-1]
			return slice[:len(slice)-1]
		}
	}
	return slice
}

func (lib *Library) CreateMember(name string) *models.Member {
	member := &models.Member{
		Id:            lib.NextMemberID,
		Name:          name,
		BorrowedBooks: []*models.Book{},
	}
	lib.Members[lib.NextMemberID] = member
	lib.NextMemberID++
	return member
}
func (lib *Library) ListMembers() []*models.Member {
	members := make([]*models.Member, 0, len(lib.Members))
	for _, m := range lib.Members {
		members = append(members, m)
	}
	return members
}
