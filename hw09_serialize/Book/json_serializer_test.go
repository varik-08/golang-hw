package book

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBook_MarshalJSON(t *testing.T) {
	book := Book{
		ID:     1,
		Title:  "The Catcher in the Rye",
		Author: "J.D. Salinger",
		Year:   1951,
		Size:   182,
		Rate:   4.3,
	}
	expectedJSON := `{"id":1,"title":"The Catcher in the Rye","author":"J.D. Salinger","year":1951,"size":182,"rate":4.3}`

	jsonData, err := json.Marshal(book)

	require.Nil(t, err)

	require.Equal(t, expectedJSON, string(jsonData))
}

func TestBook_UnmarshalJSON(t *testing.T) {
	jsonData := []byte(`{"id":1,"title":"The Catcher in the Rye","author":"J.D. Salinger","year":1951,"size":182,
"rate":4.3}`)
	var book Book
	expectedBook := Book{
		ID:     1,
		Title:  "The Catcher in the Rye",
		Author: "J.D. Salinger",
		Year:   1951,
		Size:   182,
		Rate:   4.3,
	}

	err := json.Unmarshal(jsonData, &book)

	require.Nil(t, err)

	require.Equal(t, expectedBook, book)
}
