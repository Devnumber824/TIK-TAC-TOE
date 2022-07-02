package main

import (
	"fmt"
	"pranav/game"
)

func main() {
	// var test = game.CreateNewGame("pranav", "yash")

	// test.Play(5, "pranav")
	// test.Play(4, "yash")
	// test.Play(2, "pranav")
	// test.Play(1, "yash")
	// test.Play(8, "pranav")

	// fmt.Println(test)
	count := 0
	var test *game.Game
	var name1, name2 string
	var sixe int

	for {
		fmt.Println("Please enter the player1 name:")
		_, error1 := fmt.Scanln(&name1)
		if checkError(error1) {
			fmt.Println(error1)
			fmt.Println("Please provide valid input")
			continue
		}
		fmt.Println("Please enter the player2 name:")
		_, error2 := fmt.Scanln(&name2)
		if checkError(error2) {
			fmt.Println(error2)
			fmt.Println("Please provide valid input")
			continue
		}
		fmt.Println("Please enter the sixe of the board:")
		//fmt.Scanln(&sixe)
		_, error3 := fmt.Scanln(&sixe)
		if checkError(error3) {
			fmt.Println(error3)
			fmt.Println("Please provide valid input")
			continue
		}

		test = game.CreateNewGame(name1, name2, sixe)
		fmt.Println(test)
		fmt.Println()
		for i := 0; i < len(test.MainBoard.Cells); i++ {
			if count == sixe {
				fmt.Println()
				count = 0
			}
			fmt.Print(test.MainBoard.Cells[i], " ")
			count++
		}
		fmt.Println()

		for {
			if test.Turn == sixe*sixe {
				var str = test.MainBoard.CheckWineer()
				if str == "" {
					fmt.Println("nobody won")

				} else if str == "X" {
					fmt.Println(name1, " wins")

				} else {
					fmt.Println(name2, " wins")

				}

				break
			}
			if test.Turn >= (sixe + (sixe - 1)) {
				Test := test.MainBoard.CheckWineer()

				if Test == "X" {
					fmt.Println("Winner is ", name1)
					break
				} else if Test == "O" {
					fmt.Println("Winner is ", name2)
					break
				}
			}

			if test.Turn%2 != 0 {
				fmt.Println("It's ", name1, " turn please enter Cell for play")
				var Cellno int
				_, err := fmt.Scan(&Cellno)
				if checkError(err) {
					fmt.Println(err)
					fmt.Println("please enter a valid input")
					continue
				}
				player1error := test.Play(Cellno, test.Players[0])
				if checkError(player1error) {
					fmt.Println(player1error)
					fmt.Println("please enter a valid input")
					continue
				}

			} else {
				fmt.Println("It's ", name2, " turn please enter Cell for play")
				var Cellno int
				_, err := fmt.Scan(&Cellno)
				if checkError(err) {
					fmt.Println(err)
					fmt.Println("please enter a valid input")
					continue
				}
				player2error := test.Play(Cellno, test.Players[1])
				if checkError(player2error) {
					fmt.Println(player2error)
					fmt.Println("please enter a valid input")
					continue
				}

			}

		}
		fmt.Println()
		fmt.Println("IF Want to play more press Y")
		var req string
		fmt.Scan(&req)
		if req == "Y" {
			continue
		}
		break
	}

}

func checkError(er error) bool {
	return er != nil
}
