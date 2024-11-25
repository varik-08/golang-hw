package main

import "fmt"

func main() {
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

	comparator1 := Comparator{}
	_, err := comparator1.Compare(book1, book2)
	if err != nil {
		fmt.Println("Comparator1: Error comparing books:", err)
	}

	comparator2 := NewComparator(Year)
	result, err := comparator2.Compare(book1, book2)
	if err != nil {
		fmt.Println("Error comparing books:", err)
	}
	fmt.Printf("Comparator2 returned = %v\n", result)

	comparator3 := NewComparator(Size)
	result, err = comparator3.Compare(book1, book2)
	if err != nil {
		fmt.Println("Error comparing books:", err)
	}
	fmt.Printf("Comparator3 returned = %v\n", result)

	comparator4 := NewComparator(Rate)
	result, err = comparator4.Compare(book1, book2)
	if err != nil {
		fmt.Println("Error comparing books:", err)
	}
	fmt.Printf("Comparator4 returned = %v\n", result)
}
