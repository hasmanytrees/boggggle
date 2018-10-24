package boggle

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	ClassicDice = []Die{
		{"A", "A", "C", "I", "O", "T"},
		{"A", "B", "I", "L", "T", "Y"},
		{"A", "B", "J", "M", "O", "Qu"},
		{"A", "C", "D", "E", "M", "P"},
		{"A", "C", "E", "L", "R", "S"},
		{"A", "D", "E", "N", "V", "Z"},
		{"A", "H", "M", "O", "R", "S"},
		{"B", "I", "F", "O", "R", "X"},
		{"D", "E", "N", "O", "S", "W"},
		{"D", "K", "N", "O", "T", "U"},
		{"E", "E", "F", "H", "I", "Y"},
		{"E", "G", "K", "L", "U", "Y"},
		{"E", "G", "I", "N", "T", "V"},
		{"E", "H", "I", "N", "P", "S"},
		{"E", "L", "P", "S", "T", "U"},
		{"G", "I", "L", "R", "U", "W"},
	}
)

type Game struct {
	Board [4][4]string
	Dice  []Die
	Rand  Randomizer
}

func NewGame(dice []Die) *Game {
	game := &Game{
		Board: [4][4]string{},
		Dice:  dice,
		Rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	// populate the board with die values
	numDice := len(game.Dice)
	j, k := 0, 0

	// roll each dice and put its value onto the board
	for i := 0; i < numDice; i++ {
		randomDieIndex := game.Rand.Intn(len(game.Dice))
		randomDie := game.Dice[randomDieIndex]

		game.Board[j][k] = randomDie.Roll(game.Rand)

		j = (j + 1) % 4

		if j == 0 {
			k = (k + 1) % 4
		}

		// remove the die that was rolled from the available dice so it is used only once
		game.Dice = append(game.Dice[:randomDieIndex], game.Dice[randomDieIndex+1:]...)
	}

	//game.Print()

	return game
}

func (g *Game) Print() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(g.Board[j][i] + " ")

			for k := 0; k < 3-len(g.Board[j][i]); k++ {
				fmt.Print(" ")
			}
		}

		fmt.Println("")
	}
}
