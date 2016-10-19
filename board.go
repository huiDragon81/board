package board

import "fmt"

type Board struct {
	data map[string]string
}

func (board *Board) Print() {
	for key, value := range board.data {
		fmt.Println("   " + key + " : " + value)
	}
}
