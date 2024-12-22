package book

import (
	"github.com/varik-08/golang-hw/hw09_serialize/Book/bookpb"
	"google.golang.org/protobuf/proto"
)

func SerializeBooks(books []*bookpb.Book) ([]byte, error) {
	bookList := &bookpb.BookList{Books: books}
	return proto.Marshal(bookList)
}

func DeserializeBooks(data []byte) ([]*bookpb.Book, error) {
	var bookList bookpb.BookList
	if err := proto.Unmarshal(data, &bookList); err != nil {
		return nil, err
	}
	return bookList.Books, nil
}
