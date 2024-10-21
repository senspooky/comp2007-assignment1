package expressivity

import (
	"fmt"
	"strings"

	"github.com/senspooky/comp2007-assignment1/demos"
	"github.com/senspooky/comp2007-assignment1/formatter"
)

type expressivity struct{}

func ExpressivityDemo() demos.Demo {
	return &expressivity{}
}

func (d *expressivity) Demo() {

	main()

}

type Book struct {
	Title  string
	isRead bool
}

type BookList []Book

func (b BookList) String() string {
	var sb strings.Builder
	for i, b := range b {
		var readString string
		switch b.isRead {
		case true:
			readString = "have"
		default:
			readString = "have not"
		}
		if i != 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(fmt.Sprintf("I %s read %s", readString, b.Title))
	}
	return sb.String()
}

func (b *BookList) UpdateAll(f func(Book) Book) {
	newBooks := BookList{}
	for _, b := range *b {
		b.isRead = true
		newBooks = append(newBooks, b)
	}
	*b = newBooks
}

// Go is not a very expressive language, especially when it comes to slices
func main() {
	fmt.Println(formatter.F().Format("EXPRESSIVITY DEMO"))

	// Lets make an array of books, and mark them as read
	books := BookList{
		{
			Title:  "The Infinite and The Divine",
			isRead: false,
		},
		{
			Title:  "Gaunt's Ghosts",
			isRead: false,
		},
		{
			Title:  "House of Leaves",
			isRead: false,
		},
	}

	// Print the books
	fmt.Println(books.String())
	fmt.Println()

	// mark all books as read
	fmt.Println("Marking all books as read...")
	fmt.Println()
	books.UpdateAll(func(b Book) Book {
		b.isRead = true
		return b
	})

	// Print the books
	fmt.Println("Book list: ")

	fmt.Println(books.String())

	fmt.Println()
}
