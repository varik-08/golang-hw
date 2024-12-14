package hw03

const (
	blackCell = '#'
	whiteCell = ' '
)

func generateBoard(size int) [][]rune {
	board := make([][]rune, size)

	var cell rune
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			cell = whiteCell
		} else {
			cell = blackCell
		}

		for j := 0; j < size; j++ {
			board[i] = append(board[i], cell)

			if cell == whiteCell {
				cell = blackCell
			} else {
				cell = whiteCell
			}
		}
	}

	return board
}
