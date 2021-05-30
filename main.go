package main

import (
	"fmt"
	"math/rand"
	"time"

	term "github.com/nsf/termbox-go"
)

func reset() {
	term.Sync()
}

func printPattern(newPattern [6][8]string) {
	for i := 0; i < 6; i++ {
		for j := 0; j < 8; j++ {
			fmt.Printf("%s", newPattern[i][j])
		}
		fmt.Printf("\n")
	}
}

func main() {

	locArray := [6][8]string{
		{"#", "#", "#", "#", "#", "#", "#", "#"},
		{"#", ".", ".", ".", ".", ".", ".", "#"},
		{"#", ".", "#", "#", "#", ".", ".", "#"},
		{"#", ".", ".", ".", "#", ".", "#", "#"},
		{"#", ".", "#", ".", ".", ".", ".", "#"},
		{"#", "#", "#", "#", "#", "#", "#", "#"},
	}

	treasureProbablePos := [17][2]int{
		{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6},
		{2, 1}, {2, 5}, {2, 6},
		{3, 1}, {3, 2}, {3, 3}, {3, 5},
		{4, 3}, {4, 4}, {4, 5}, {4, 6},
	}

	// treasure to random position on the map
	rand.Seed(time.Now().UnixNano())
	treasurePos := rand.Intn(16)
	rowTreasure := treasureProbablePos[treasurePos][0]
	colTreasure := treasureProbablePos[treasurePos][1]

	locArray[rowTreasure][colTreasure] = "$"

	rowMin := 1
	rowMax := 4
	columnMax := 6

	newPattern := locArray

	// player's start position
	row := 4
	column := 1
	newPattern[row][column] = "X"

	err := term.Init()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 17; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("%d ", treasureProbablePos[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Println("====================================================================\n")

	printPattern(newPattern)

	fmt.Println("Press Up, Right, Down key to move. Press ESC button to quit")

	lastMove := "start"

keyPressListenerLoop:
	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				break keyPressListenerLoop
			case term.KeyArrowUp:
				if lastMove == "start" || lastMove == "up" {
					newPattern := locArray
					if row-1 >= rowMin {
						row -= 1
						if newPattern[row][column] != "#" {
							newPattern[row][column] = "X"
							printPattern(newPattern)

							if newPattern[rowTreasure][colTreasure] == "X" {
								fmt.Println("Treasure Found, Congratulation!")
								break keyPressListenerLoop
							}
							reset()
							fmt.Println("Move up")
							lastMove = "up"
						} else {
							row += 1
							newPattern[row][column] = "X"
							printPattern(newPattern)

							fmt.Println("Path is Blocked, try another move")
						}
					} else {
						fmt.Println("Path is Blocked, try another move")
					}
				} else {
					fmt.Println("Path is Blocked, try another move")
				}

			case term.KeyArrowRight:
				if lastMove == "up" || lastMove == "right" {
					newPattern := locArray
					if column+1 <= columnMax {

						column += 1
						if newPattern[row][column] != "#" {
							newPattern[row][column] = "X"
							printPattern(newPattern)

							// if new position is the location of the treasure. game end and player win
							if newPattern[rowTreasure][colTreasure] == "X" {
								fmt.Println("Treasure Found, Congratulation!")
								break keyPressListenerLoop
							}
							reset()
							fmt.Println("Move Right")
							lastMove = "right"
						} else {
							column -= 1
							newPattern[row][column] = "X"
							printPattern(newPattern)

							fmt.Println("Path is Blocked, try another move")
						}
					} else {
						fmt.Println("Path is Blocked, try another move")
					}
				} else {
					fmt.Println("Can't move to right direction")
				}

			case term.KeyArrowDown:
				if lastMove == "right" || lastMove == "down" {
					newPattern := locArray
					if row+1 <= rowMax {
						row += 1
						if newPattern[row][column] != "#" {
							newPattern[row][column] = "X"
							printPattern(newPattern)

							if newPattern[rowTreasure][colTreasure] == "X" {
								fmt.Println("Treasure Found, Congratulation!")
								break keyPressListenerLoop
							}
							reset()
							fmt.Println("Move Down")
							lastMove = "down"
						} else {
							row -= 1
							newPattern[row][column] = "X"
							printPattern(newPattern)
							if newPattern[rowTreasure][colTreasure] == "X" {
								fmt.Println("Treasure Found, Congratulation!")
							}
							fmt.Println("Path is Blocked, try another move")
						}
					} else {
						fmt.Println("Path is Blocked, try another move")
					}
				} else {
					fmt.Println("Can't move down!")
				}

			default:
				reset()
				fmt.Println("Only Up, Right, or Down key")

			}
		case term.EventError:
			panic(ev.Err)
		}
	}
}
