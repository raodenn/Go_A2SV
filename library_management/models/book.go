package models

type Book struct {
	Id     int
	Title  string
	Author string
	Status string
}

var DummyBooks = []Book{
	{Id: 1, Title: "To Kill a Mockingbird", Author: "Harper Lee", Status: "Available"},
	{Id: 2, Title: "1984", Author: "George Orwell", Status: "Available"},
	{Id: 3, Title: "Mistborn: The Final Empire", Author: "Brandon Sanderson", Status: "Available"},
	{Id: 4, Title: "Pride and Prejudice", Author: "Jane Austen", Status: "Available"},
	{Id: 5, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Status: "Available"},
	{Id: 6, Title: "Warbreaker", Author: "Brandon Sanderson", Status: "Available"},
	{Id: 7, Title: "The Way of Kings", Author: "Brandon Sanderson", Status: "Available"},
	{Id: 8, Title: "The Hobbit", Author: "J.R.R. Tolkien", Status: "Available"},
	{Id: 9, Title: "Fahrenheit 451", Author: "Ray Bradbury", Status: "Available"},
	{Id: 10, Title: "Oathbringer", Author: "Brandon Sanderson", Status: "Available"},
	{Id: 11, Title: "Jane Eyre", Author: "Charlotte Brontë", Status: "Available"},
}
