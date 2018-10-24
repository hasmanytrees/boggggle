package boggle

import (
	"fmt"
	"testing"
)

var (
	TestBoard = [4][4]string{
		{"P", "C", "A", "L"},
		{"F", "N", "I", "M"},
		{"D", "J", "G", "B"},
		{"E", "O", "H", "K"},
	}
)

func TestPrint(t *testing.T) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(TestBoard[j][i] + " ")

			for k := 0; k < 3-len(TestBoard[j][i]); k++ {
				fmt.Print(" ")
			}
		}

		fmt.Println("")
	}
}

func TestGet(t *testing.T) {
	grader := Grader{
		Board:    TestBoard,
		Position: Position{0, 0},
	}

	value, ok := grader.Get(-1, -1)

	if value != "" {
		t.Fatalf("There should not be a NW neighbor for die at position 0,0 ... value should be empty but got %s", value)
	}

	if ok {
		t.Fatalf("There should not be a NW neighbor for die at position 0,0 ... ok should be false")
	}
}

func TestMove(t *testing.T) {
	grader := Grader{
		Board:    TestBoard,
		Position: Position{0, 0},
	}

	ok := grader.Move(-1, -1)

	if ok {
		t.Fatalf("You should not be able to move W 0,0 ... ok should be false")
	}
}

func TestWordContainsPosition(t *testing.T) {
	grader := &Grader{
		Board:    TestBoard,
		Position: Position{2, 0}, // Should be D on the TestBoard
		Path:     []Position{},
	}

	grader.Move(1, 1) // Should be O

	positionFound := false

	for _, v := range grader.Path {
		if v.X == 3 && v.Y == 1 {
			positionFound = true
			break
		}
	}

	if !positionFound {
		t.Fatalf("When moving to a new position the position needs to be added to the Path, expected (4, 1) to be in the path %v", grader.Path)
	}
}

func TestSliceWord(t *testing.T) {
	word := "QUICK"

	got := SliceWord(word)

	fmt.Println(got)

	if got[0] != "Qu" {
		t.Fatalf("Expected %s to be Qu", got[0])
	}
}

func TestGrader_WordExists(t *testing.T) {
	grader := &Grader{
		Board:    TestBoard,
		Position: Position{0, 0}, // Should be D on the TestBoard
		Path:     []Position{},
	}

	path, ok := grader.WordExists("dog")

	fmt.Printf("Path for dog: %v\n", path)

	if !ok {
		t.Fatalf("Expected %s to be found", "dog")
	}

	path, ok = grader.WordExists("hog")
	fmt.Printf("Path for hog: %v\n", path)

	if !ok {
		t.Fatalf("Expected %s to be found", "hog")
	}

	path, ok = grader.WordExists("mind")
	fmt.Printf("Path for mind: %v\n", path)

	if !ok {
		t.Fatalf("Expected %s to be found", "mind")
	}

	path, ok = grader.WordExists("lamb")
	fmt.Printf("Path for lamb: %v\n", path)

	if !ok {
		t.Fatalf("Expected %s to be found", "lamb")
	}

	path, ok = grader.WordExists("cage")
	fmt.Printf("Path for cage: %v\n", path)

	if ok {
		t.Fatalf("Expected %s to NOT be found", "cage")
	}
}

//func TestFindWord(t *testing.T){
//	grader := &Grader{
//		Board: TestBoard,
//		Position: Position{0, 0},// Should be D on the TestBoard
//		Path: []Position{},
//	}
//
//	path, ok := grader.FindWord(SliceWord("dog"), Position{2,0}, []Position{})
//
//	fmt.Printf("Path for dog: %v\n", path)
//
//	if !ok {
//		t.Fatalf("Expected %s to be found", "dog")
//	}
//
//
//	path, ok = grader.FindWord(SliceWord("hog"), Position{3,2}, []Position{})
//
//	fmt.Printf("Path for hog: %v\n", path)
//
//	if !ok {
//		t.Fatalf("Expected %s to be found", "hog")
//	}
//
//
//	path, ok = grader.FindWord(SliceWord("mind"), Position{1,3}, []Position{})
//
//	fmt.Printf("Path for mind: %v\n", path)
//
//	if !ok {
//		t.Fatalf("Expected %s to be found", "mind")
//	}
//}
