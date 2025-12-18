package main

import (
	"fmt"

	"github.com/Ekosik7/Assignment1/Bank"
	"github.com/Ekosik7/Assignment1/Company"
	"github.com/Ekosik7/Assignment1/Library"
	"github.com/Ekosik7/Assignment1/Shapes"
)

func main() {
	lib := Library.NewLibrary()

	lib.AddBook(Library.Book{
		ID:     "1",
		Title:  "Go Basics",
		Author: "Alan Donovan",
	})
	lib.AddBook(Library.Book{
		ID:     "2",
		Title:  "Clean Code",
		Author: "Robert Martin",
	})

	lib.BorrowBook("1")

	for _, b := range lib.ListAvailableBooks() {
		fmt.Println(b.Title)
	}

	shapes := []Shapes.Shape{
		Shapes.Rectangle{Width: 4, Height: 5},
		Shapes.Circle{Radius: 3},
		Shapes.Square{Side: 2},
	}

	for _, s := range shapes {
		fmt.Printf("%.2f %.2f\n", s.Area(), s.Perimeter())
	}

	company := Company.NewCompany()
	company.AddEmployee(1, Company.FullTimeEmployee{Name: "Alice", Salary: 3000})
	company.AddEmployee(2, Company.PartTimeEmployee{Name: "Bob", HourlyPay: 10, Hours: 80})

	for _, info := range company.ListEmployees() {
		fmt.Println(info)
	}

	account := Bank.NewBankAccount("ACC1", "John")
	account.Deposit(1000)
	account.Withdraw(250)

	fmt.Println(account.GetBalance())
}
