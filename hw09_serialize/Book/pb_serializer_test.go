package book

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/varik-08/golang-hw/hw09_serialize/Book/bookpb"
)

func TestSerializeBooksSuccess(t *testing.T) {
	books := []*bookpb.Book{
		{Id: 1, Title: "Book 1", Author: "Author 1", Year: 2022, Size: 100, Rate: 4.5},
		{Id: 2, Title: "Book 2", Author: "Author 2", Year: 2021, Size: 150, Rate: 4.3},
	}

	serializedData, err := SerializeBooks(books)

	require.Nil(t, err)
	require.NotNil(t, serializedData)
}

func TestDeserializeBooks(t *testing.T) {
	books := []*bookpb.Book{
		{Id: 1, Title: "Book 1", Author: "Author 1", Year: 2022, Size: 100, Rate: 4.5},
		{Id: 2, Title: "Book 2", Author: "Author 2", Year: 2021, Size: 150, Rate: 4.3},
	}

	serializedData, err := SerializeBooks(books)

	require.Nil(t, err)

	deserializedBooks, err := DeserializeBooks(serializedData)

	require.Nil(t, err)

	for i, book := range deserializedBooks {
		require.Equal(t, books[i].Id, book.Id)
		require.Equal(t, books[i].Title, book.Title)
		require.Equal(t, books[i].Author, book.Author)
		require.Equal(t, books[i].Year, book.Year)
		require.Equal(t, books[i].Size, book.Size)
		require.Equal(t, books[i].Rate, book.Rate)
	}
}
