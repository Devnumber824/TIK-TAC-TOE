package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"pranav/board"
	"pranav/game"
	"strconv"
	"strings"
)

func main() {
	//count := 0

	var test *game.Game
	var name1, name2 string
	var sym1error, sym2error, cellerror1, cellerror2 error
	var testinga string
	var sixe int
	var sym1, sym2 string
	var error1, error2, error3 error
	var cell string
	var cellno int
	reader := bufio.NewReader(os.Stdin)
	//var t = bufio.NewReader(os.Stdin)
	for {
		for {
			fmt.Println("Please enter the player1 name:")
			name1, error1 = reader.ReadString('\n')
			name1 = strings.TrimSuffix(name1, "\r\n")
			//name1, error1 = t.ReadString('\n')
			if checkError(error1) {
				fmt.Println(error1)
				fmt.Println("Please provide valid input")
				continue
			} else {
				break
			}
		}
		for {
			fmt.Println("Please enter the player1 Symbol:")
			sym1, sym1error = reader.ReadString('\n')
			sym1 = strings.TrimSuffix(sym1, "\r\n")
			if len(sym1) > 1 {
				sym1error = errors.New("enter a single character")
			}
			if checkError(sym1error) {
				fmt.Println(sym1error)
				fmt.Println("Please provide valid input", len(sym1))
				continue
			} else {
				break
			}
		}
		for {
			fmt.Println("Please enter the player2 name:")
			name2, error2 = reader.ReadString('\n')
			name2 = strings.TrimSuffix(name2, "\r\n")
			if name2 == name1 {
				error2 = errors.New("you entered same name for player 1 and 2")
			}
			if checkError(error2) {
				fmt.Println(error2)
				fmt.Println("Please provide valid input")
				continue
			} else {
				break
			}
		}
		for {
			fmt.Println("Please enter the player2 Symbol:")
			sym2, sym2error = reader.ReadString('\n')
			sym2 = strings.TrimSuffix(sym2, "\r\n")
			if len(sym2) != 1 {
				sym2error = errors.New("enter a single character")
			}
			if sym1 == sym2 {
				sym2error = errors.New("symbol is same")
			}
			if checkError(sym2error) {
				fmt.Println(sym2error)
				fmt.Println("Please provide valid input")
				continue
			} else {
				break
			}
		}

		for {
			fmt.Println("Please enter the sixe of the board b/w 3 to 7:")
			testinga, error3 = reader.ReadString('\n')
			testinga = strings.TrimSuffix(testinga, "\r\n")
			//_, error3 = fmt.Scan(&testinga)
			if len(testinga) > 1 {
				error3 = errors.New("wrong input")
			} else {

				sixe, error3 = strconv.Atoi(testinga)
			}
			if sixe < 3 || sixe > 7 {
				error3 = errors.New("please enter value b/w 3 and 7")
			}
			if checkError(error3) {
				fmt.Println(error3)
				fmt.Println("Please provide valid input", sixe, testinga, len(testinga))
				continue
			} else {
				break

			}
		}

		test = game.CreateNewGame(name1, name2, sixe, sym1, sym2)
		fmt.Println(test)
		fmt.Println()
		test.MainBoard.Printboard()

		fmt.Println()

		for {
			if test.Turn == sixe*sixe+1 {
				var str, ans = test.MainBoard.CheckWineer()
				if str == "" {
					fmt.Println("nobody won ", ans)
					test.MainBoard.Printboard()

				} else if str == sym1 {
					fmt.Println(name1, " wins ", ans)
					test.MainBoard.Printboard()
				} else {
					fmt.Println(name2, " wins ", ans)
					test.MainBoard.Printboard()

				}

				break
			}
			if test.Turn >= (sixe + (sixe - 1)) {
				Test, ans := test.MainBoard.CheckWineer()

				if Test == sym1 {
					fmt.Println("Winner is ", name1, " ", ans)
					test.MainBoard.Printboard()
					break
				} else if Test == sym2 {
					fmt.Println("Winner is ", name2, " ", ans)
					test.MainBoard.Printboard()
					break
				}
			}

			if test.Turn%2 != 0 {
				fmt.Println("It's ", name1, " turn please enter Cell for play")
				// var Cellno int
				// _, err := fmt.Scan(&Cellno)
				cell, cellerror1 = reader.ReadString('\n')
				cell = strings.TrimSuffix(cell, "\r\n")
				if len(cell) != 1 {
					cellerror1 = errors.New("wrong value")
				} else {
					cellno, cellerror1 = strconv.Atoi(cell)
				}

				if cellno < 0 || cellno >= sixe*sixe {
					cellerror1 = errors.New("value extends the limits")
				}

				if checkError(cellerror1) {
					fmt.Println(cellerror1)
					fmt.Println("please enter a valid input")
					continue
				}
				player1error := test.Play(cellno, test.Players[0])
				if checkError(player1error) {
					fmt.Println(player1error)
					fmt.Println("please enter a valid input")
					continue
				}
				test.MainBoard.Printboard()

			} else {
				fmt.Println("It's ", name2, " turn please enter Cell for play")
				// var Cellno int
				// _, err := fmt.Scanln(&Cellno)
				cell, cellerror2 = reader.ReadString('\n')
				cell = strings.TrimSuffix(cell, "\r\n")
				if len(cell) != 1 {
					cellerror2 = errors.New("wrong value")
				} else {
					cellno, cellerror2 = strconv.Atoi(cell)
				}

				if cellno < 0 || cellno >= sixe*sixe {
					cellerror2 = errors.New("value extends the limits")
				}

				if checkError(cellerror2) {
					fmt.Println(cellerror2)
					fmt.Println("please enter a valid input")
					continue
				}
				player2error := test.Play(cellno, test.Players[1])
				if checkError(player2error) {
					fmt.Println(player2error)
					fmt.Println("please enter a valid input")
					continue
				}
				test.MainBoard.Printboard()
			}

		}
		fmt.Println()
		fmt.Println("IF Want to play more press Y")
		var req string
		fmt.Scan(&req)
		if req == "Y" {
			board.MainBoard = board.Nilboard

			continue
		}
		break
	}

}

func checkError(er error) bool {
	return er != nil
}
