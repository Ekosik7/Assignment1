package Library

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func NewLibrary() *Library {
	return &Library{
		Books: make(map[string]Book),
	}
}

func (l *Library) AddBook(book Book) {
	l.Books[book.ID] = book
}

func (l *Library) BorrowBook(id string) bool {
	book, ok := l.Books[id]
	if !ok || book.IsBorrowed {
		return false
	}
	book.IsBorrowed = true
	l.Books[id] = book
	return true
}

func (l *Library) ReturnBook(id string) bool {
	book, ok := l.Books[id]
	if !ok || !book.IsBorrowed {
		return false
	}
	book.IsBorrowed = false
	l.Books[id] = book
	return true
}

func (l *Library) ListAvailableBooks() []Book {
	var result []Book
	for _, book := range l.Books {
		if !book.IsBorrowed {
			result = append(result, book)
		}
	}
	return result
}
