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

	lib.AddBook(Library.Book{ID: "1", Title: "Harry Potter and the Philosopher's Stone", Author: "J.K. Rowling"})
	lib.AddBook(Library.Book{ID: "2", Title: "The Hobbit", Author: "J.R.R. Tolkien"})
	lib.AddBook(Library.Book{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy"})
	lib.AddBook(Library.Book{ID: "4", Title: "Pride and Prejudice", Author: "Jane Austen"})

	// Borrow one example book
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
	company.AddEmployee(1, Company.FullTimeEmployee{Name: "Nursultan", Salary: 3200})
	company.AddEmployee(2, Company.FullTimeEmployee{Name: "Aigerim", Salary: 2800})
	company.AddEmployee(3, Company.PartTimeEmployee{Name: "Almas", HourlyPay: 12, Hours: 40})
	company.AddEmployee(4, Company.PartTimeEmployee{Name: "Zhansaya", HourlyPay: 9, Hours: 60})

	for _, info := range company.ListEmployees() {
		fmt.Println(info)
	}

	account := Bank.NewBankAccount("ACC1", "John")
	account.Deposit(1000)
	account.Withdraw(250)

	fmt.Println(account.GetBalance())
}
