package boggle

import "testing"

var (
	TestDice = []Die{
		{"A", "A", "A", "A", "A", "A"},
		{"B", "B", "B", "B", "B", "B"},
		{"C", "C", "C", "C", "C", "C"},
		{"D", "D", "D", "D", "D", "D"},
		{"E", "E", "E", "E", "E", "E"},
		{"F", "F", "F", "F", "F", "F"},
		{"G", "G", "G", "G", "G", "G"},
		{"H", "H", "H", "H", "H", "H"},
		{"I", "I", "I", "I", "I", "I"},
		{"J", "J", "J", "J", "J", "J"},
		{"K", "K", "K", "K", "K", "K"},
		{"L", "L", "L", "L", "L", "L"},
		{"M", "M", "M", "M", "M", "M"},
		{"N", "N", "N", "N", "N", "N"},
		{"O", "O", "O", "O", "O", "O"},
		{"P", "P", "P", "P", "P", "P"},
	}
)

func TestNewGame(t *testing.T) {
	//setup
	game := NewGame(TestDice)

	//exercise

	//assert
	gotBoardWidth := len(game.Board)
	gotBoardHeight := len(game.Board[0])

	// check board width
	if gotBoardWidth != 4 {
		t.Fatalf("Width is not correct, wanted 4 and got %d", gotBoardWidth)
	}

	// check board height
	if gotBoardHeight != 4 {
		t.Fatalf("Height is not correct, wanted 4 and got %d", gotBoardHeight)
	}

	// check board is populated with non-empty strings
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if game.Board[i][j] == "" {
				t.Fatalf("Board is not populated with die values")
			}
		}
	}

	// make sure board is populated with one of every letter (A-P)
	foundDie := map[string]bool{
		"A": false,
		"B": false,
		"C": false,
		"D": false,
		"E": false,
		"F": false,
		"G": false,
		"H": false,
		"I": false,
		"J": false,
		"K": false,
		"L": false,
		"M": false,
		"N": false,
		"O": false,
		"P": false,
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			foundDie[game.Board[i][j]] = true
		}
	}

	for k, v := range foundDie {
		if !v {
			t.Fatalf("Board does not contain a die from TestDie: %s", k)
		}
	}

	// check dice are empty after populating board
	if len(game.Dice) > 0 {
		t.Fatalf("Not all dice were used to populate the board")
	}
}
