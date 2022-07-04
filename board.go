package board

import (
	"fmt"
	"pranav/cell"
	"strconv"
)

type Board struct {
	Cells []cell.Cell
	Sixe  int
}

var MainBoard *Board
var Nilboard *Board

func createNewBoard(n int) *Board {
	var newb Board
	newb.Cells = make([]cell.Cell, n*n)
	newb.Sixe = n
	for i := 0; i < len(newb.Cells); i++ {
		var str string
		if i <= 9 {
			str = "0" + strconv.Itoa(i)
		} else {
			str = strconv.Itoa(i)

		}

		newb.Cells[i].Mark = str
	}
	return &newb
}

func CreateNewBoard(n int) *Board {
	if MainBoard != nil {

		return MainBoard
	}
	MainBoard = createNewBoard(n)

	return MainBoard
}

func (b *Board) CheckWineer() (string, string) {
	verwin, str1 := b.checkVertical()
	if verwin {
		var ans = "this is vertical win"
		return str1, ans
	}
	horwin, str2 := b.checkHor()
	if horwin {
		var ans1 = "this is hor win"
		return str2, ans1
	}
	dwin, str3 := b.diagonalCheck()
	if dwin {
		var ans2 = "this is dig win"
		return str3, ans2
	}
	return "", ""

}

func (b *Board) checkVertical() (bool, string) {
	var vercount1 int
	var vercount2 int
	for i := 0; i < b.Sixe; i++ {
		for j := i; j < b.Sixe*b.Sixe; j = j + b.Sixe {
			if b.Cells[j].Mark == "X" {
				vercount1++
			} else if b.Cells[j].Mark == "O" {
				vercount2++
			}
		}
		if vercount1 == b.Sixe {
			return true, "X"
		} else if vercount2 == b.Sixe {
			return true, "O"
		}
		vercount1 = 0
		vercount2 = 0
	}
	return false, ""
}

func (b *Board) checkHor() (bool, string) {
	var horcount1 int
	var horcount2 int
	count := 1
	for i := 0; i < b.Sixe*b.Sixe; i = i + b.Sixe {
		for j := i; count <= b.Sixe; j++ {
			if b.Cells[j].Mark == "X" {
				horcount1++
			} else if b.Cells[j].Mark == "O" {
				horcount2++
			}
			count++
		}
		if horcount1 == b.Sixe {
			return true, "X"
		} else if horcount2 == b.Sixe {
			return true, "O"
		}
		horcount1 = 0
		horcount2 = 0
		count = 1
	}
	return false, ""
}

func (b *Board) diagonalCheck() (bool, string) {
	var digcount1 int
	var digcount2 int

	for j := 0; j < b.Sixe*b.Sixe; j += b.Sixe + 1 {
		if b.Cells[j].Mark == "X" {
			digcount1++
		} else if b.Cells[j].Mark == "O" {
			digcount2++
		}
	}
	if digcount1 == b.Sixe {
		return true, "X"
	} else if digcount2 == b.Sixe {
		return true, "O"
	}
	count1 := 0
	count2 := 0
	for j := b.Sixe - 1; j < b.Sixe*b.Sixe-1; j += b.Sixe - 1 {
		if b.Cells[j].Mark == "X" {
			count1++
		} else if b.Cells[j].Mark == "O" {
			count2++
		}
	}
	if count1 == b.Sixe {
		return true, "X"
	} else if count2 == b.Sixe {
		return true, "O"
	}

	return false, ""
}

func (b *Board) Printboard() {
	count := 0
	for i := 0; i < len(b.Cells); i++ {
		if count == b.Sixe {
			fmt.Println()
			count = 0
		}
		fmt.Print(b.Cells[i], " ")
		count++
	}
	fmt.Println()

}

// for i := 0; i < 8; i++ {
// 	line := ""
// 	switch i {
// 	case 0:
// 		line = b.Cells[0].Mark + b.Cells[1].Mark + b.Cells[2].Mark

// 	case 1:
// 		line = b.Cells[0].Mark + b.Cells[3].Mark + b.Cells[6].Mark

// 	case 2:
// 		line = b.Cells[0].Mark + b.Cells[4].Mark + b.Cells[8].Mark

// 	case 3:
// 		line = b.Cells[3].Mark + b.Cells[4].Mark + b.Cells[5].Mark

// 	case 4:
// 		line = b.Cells[6].Mark + b.Cells[7].Mark + b.Cells[8].Mark

// 	case 5:
// 		line = b.Cells[2].Mark + b.Cells[4].Mark + b.Cells[6].Mark

// 	case 6:
// 		line = b.Cells[1].Mark + b.Cells[4].Mark + b.Cells[7].Mark

// 	case 7:
// 		line = b.Cells[2].Mark + b.Cells[5].Mark + b.Cells[8].Mark

// 	}
// 	if line == "XXX" {
// 		return "X"
// 	} else if line == "OOO" {
// 		return "O"
// 	}

// }
// return "m"
