package board

import (
	"fmt"
	"pranav/cell"
	"testing"
)

func TestCheckwin1(t *testing.T) {
	var test Board = Board{
		Cells: make([]cell.Cell, 8),
		Sixe:  3,
	}

	test.Cells[0].Mark = "X"
	test.Cells[1].Mark = "X"
	test.Cells[2].Mark = "O"
	test.Cells[3].Mark = "X"
	test.Cells[4].Mark = "O"
	test.Cells[6].Mark = "X"
	test.Cells[7].Mark = "X"

	str1, str2 := test.CheckWineer()

	if str1 == "X" {
		fmt.Println("Test 1 is passed", str2)
	} else {
		t.Error("Test failed")
	}

}

func TestCheckwin2(t *testing.T) {
	var test1 Board = Board{
		Cells: make([]cell.Cell, 16),
		Sixe:  4,
	}
	//	fmt.Println(test1.Cells)
	test1.Cells[1].Mark = "X"
	test1.Cells[2].Mark = "O"
	test1.Cells[3].Mark = "X"
	test1.Cells[4].Mark = "O"
	test1.Cells[6].Mark = "X"
	test1.Cells[7].Mark = "X"
	test1.Cells[8].Mark = "O"
	test1.Cells[9].Mark = "X"
	test1.Cells[11].Mark = "O"
	test1.Cells[12].Mark = "X"

	str1, str2 := test1.CheckWineer()

	if str1 == "X" {
		fmt.Println("Test 2 is passed", str2)
	} else {
		t.Error("Test failed")
	}

}

func TestCheckwin3(t *testing.T) {
	var test Board = Board{
		Cells: make([]cell.Cell, 25),
		Sixe:  5,
	}

	test.Cells[0].Mark = "X"
	test.Cells[1].Mark = "X"
	test.Cells[2].Mark = "O"
	test.Cells[3].Mark = "X"
	test.Cells[4].Mark = "O"
	test.Cells[5].Mark = "O"
	test.Cells[6].Mark = "X"
	test.Cells[7].Mark = "X"
	test.Cells[8].Mark = "O"
	test.Cells[9].Mark = "X"
	test.Cells[10].Mark = "O"
	test.Cells[11].Mark = "X"
	test.Cells[12].Mark = "X"
	test.Cells[13].Mark = "O"
	test.Cells[14].Mark = "O"
	test.Cells[15].Mark = "X"
	test.Cells[16].Mark = "O"
	test.Cells[17].Mark = "O"
	test.Cells[18].Mark = "X"
	test.Cells[19].Mark = "O"
	test.Cells[20].Mark = "O"
	test.Cells[21].Mark = "X"
	test.Cells[22].Mark = "O"
	test.Cells[23].Mark = "O"
	test.Cells[24].Mark = "X"

	str1, str2 := test.CheckWineer()

	if str1 == "X" {
		fmt.Println("Test 3 is passed", str2)
	} else {
		t.Error("Test failed")
	}

}

func TestCheckwin4(t *testing.T) {
	var test Board = Board{
		Cells: make([]cell.Cell, 25),
		Sixe:  5,
	}

	test.Cells[0].Mark = "X"
	test.Cells[1].Mark = "X"
	test.Cells[2].Mark = "O"
	test.Cells[3].Mark = "X"
	test.Cells[4].Mark = "O"
	test.Cells[5].Mark = "O"
	test.Cells[6].Mark = "X"
	test.Cells[7].Mark = "X"
	test.Cells[8].Mark = "O"
	test.Cells[9].Mark = "X"
	test.Cells[10].Mark = "O"
	test.Cells[11].Mark = "X"
	test.Cells[12].Mark = "X"
	test.Cells[13].Mark = "O"
	test.Cells[14].Mark = "O"
	test.Cells[15].Mark = "X"
	test.Cells[16].Mark = "O"
	test.Cells[17].Mark = "O"
	test.Cells[18].Mark = "X"
	test.Cells[19].Mark = "O"
	test.Cells[20].Mark = "O"
	test.Cells[21].Mark = "X"
	test.Cells[22].Mark = "O"
	test.Cells[23].Mark = "O"
	test.Cells[24].Mark = "O"

	str1, str2 := test.CheckWineer()

	if str1 == "X" || str1 == "O" {
		fmt.Println(str1, str2)
		test.Printboard()
		t.Error("Test failed")
	} else {
		fmt.Println("Testcase 4 is passed", str2)
	}

}
