package hw04_struct_comparator

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func getTestBooks() (Book, Book) {
	book1 := Book{
		id:     1,
		title:  "Title 1",
		author: "Author 1",
		year:   2010,
		size:   100,
		rate:   4.5,
	}
	book2 := Book{
		id:     1,
		title:  "Title 2",
		author: "Author 2",
		year:   2011,
		size:   50,
		rate:   5,
	}

	return book1, book2
}

func TestCompareWithError(t *testing.T) {
	book1, book2 := getTestBooks()

	comparator1 := Comparator{}
	_, err := comparator1.Compare(book1, book2)

	require.NotNil(t, err)
}

func TestCompareYear(t *testing.T) {
	book1, book2 := getTestBooks()

	comparator := NewComparator(Year)
	result, _ := comparator.Compare(book1, book2)

	require.False(t, result)
}

func TestCompareSize(t *testing.T) {
	book1, book2 := getTestBooks()

	comparator := NewComparator(Size)
	result, _ := comparator.Compare(book1, book2)

	require.True(t, result)
}

func TestCompareRate(t *testing.T) {
	book1, book2 := getTestBooks()

	comparator := NewComparator(Rate)
	result, _ := comparator.Compare(book1, book2)

	require.False(t, result)
}
