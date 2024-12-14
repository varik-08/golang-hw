package hw04_struct_comparator

import "fmt"

type Field int

const (
	Year Field = iota + 1
	Size
	Rate
)

type Comparator struct {
	field Field
}

func NewComparator(field Field) *Comparator {
	return &Comparator{field: field}
}

func (comparator *Comparator) Compare(book1 Book, book2 Book) (bool, error) {
	switch comparator.field {
	case Year:
		return book1.Year() > book2.Year(), nil
	case Size:
		return book1.Size() > book2.Size(), nil
	case Rate:
		return book1.Rate() > book2.Rate(), nil
	default:
		return false, fmt.Errorf("create comparator with NewComparator")
	}
}
