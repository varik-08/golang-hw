package main

import "fmt"

func main() {
	const blackCell = '#'
	const whiteCell = ' '
	var size int

	fmt.Print("Write size of board: ")

	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		fmt.Printf("Error read size: %v", err)
		return
	}

	var cell rune
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			cell = whiteCell
		} else {
			cell = blackCell
		}

		for j := 0; j < size; j++ {
			fmt.Print(string(cell))

			if cell == whiteCell {
				cell = blackCell
			} else {
				cell = whiteCell
			}
		}
		fmt.Println()
	}
}
