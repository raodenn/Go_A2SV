package models

type Member struct {
	Id            int
	Name          string
	BorrowedBooks []*Book
}

var DummyMembers = []Member{
	{Id: 1, Name: "Arin Lightfoot"},
	{Id: 2, Name: "Selena Starwatcher"},
	{Id: 3, Name: "Thorin Oakshield"},
	{Id: 4, Name: "Kaela Windrider"},
	{Id: 5, Name: "Brando Sando"},
}
