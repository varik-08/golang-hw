package book

import "encoding/json"

func (b *Book) MarshalJSON() ([]byte, error) {
	type BookAlias Book

	tempBook := &struct {
		*BookAlias
	}{
		BookAlias: (*BookAlias)(b),
	}

	return json.Marshal(tempBook)
}

func (b *Book) UnmarshalJSON(data []byte) error {
	type BookAlias Book

	tempBook := &struct {
		*BookAlias
	}{
		BookAlias: (*BookAlias)(b),
	}

	return json.Unmarshal(data, tempBook)
}
