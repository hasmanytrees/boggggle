package boggle

import (
	"strings"
)

type Grader struct {
	Board    [4][4]string
	Position Position
	Path     []Position
}

type Position struct {
	X int
	Y int
}

//
//func (g *Grader) WordExists(word string) bool {
//	letters := SliceWord(word)
//	//startingPositions := findStartingPositions(letters[0], g)
//		for i := 0; i < 4; i++ {
//			for j := 0; j < 4; j++ {
//				if g.Board[i][j] == letters[0] {
//					g.Position.X = i
//					g.Position.Y = i
//				}
//			}
//		}
//
//	//for _, v := range letters {
//	//	for i := 0; i < 4; i++ {
//	//		for j := 0; j < 4; j++ {
//	//				if g.Board[i][j] == v {
//	//					g.Position.X = i
//	//					g.Position.Y = i
//	//				}
//	//			}
//	//		}
//	//}
//	return false
//}

func (g *Grader) Get(xDelta int, yDelta int) (string, bool) {
	newX := g.Position.X + xDelta
	newY := g.Position.Y + yDelta

	if newX >= 0 && newX < len(g.Board[0]) && newY >= 0 && newY < len(g.Board) {
		return g.Board[newX][newY], true
	}

	return "", false
}

func (g *Grader) Move(xDelta int, yDelta int) bool {
	newX := g.Position.X + xDelta
	newY := g.Position.Y + yDelta

	if newX >= 0 && newX < len(g.Board[0]) && newY >= 0 && newY < len(g.Board) {
		g.Position.X = newX
		g.Position.Y = newY
		g.Path = append(g.Path, Position{newX, newY})
		return true
	}

	return false
}

func SliceWord(word string) []string {
	word = strings.ToUpper(word)

	var letters = []string{}

	skip := false

	for _, c := range word {
		s := string(c)

		if s == "Q" {
			letters = append(letters, "Qu")
			skip = true
		} else {
			if !skip {
				letters = append(letters, s)
			}
			skip = false
		}
	}

	return letters
}

func (g *Grader) findStartingPositions(firstLetter string) []Position {
	var positions = []Position{}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.Board[i][j] == firstLetter {
				positions = append(positions, Position{i, j})
			}
		}
	}

	return positions
}

func (g *Grader) WordExists(word string) ([]Position, bool) {
	word = strings.ToUpper(word)

	firstPositions := g.findStartingPositions(string(word[0]))

	for _, v := range firstPositions {
		path, ok := g.FindWord(SliceWord(word), v, []Position{})

		if ok {
			return path, true
		}
	}

	return nil, false
}

func (g *Grader) FindWord(s []string, pos Position, path []Position) ([]Position, bool) {
	//fmt.Printf("Looking for: %v at %v\n", s, pos)

	result := false

	// check to make sure that the location to check is valid
	if pos.X < 0 || pos.X >= len(g.Board[0]) || pos.Y < 0 || pos.Y >= len(g.Board) {
		return nil, false
	}

	// word slice length is just 1, either we found the holy grail or we didn't
	if len(s) == 1 {
		if g.Board[pos.X][pos.Y] == s[0] {
			path = append(path, pos)

			return path, true
		} else {
			return nil, false
		}
	}

	// there is more to search for but we aren't on a matching position of the first item then no dice (pun intended)
	if g.Board[pos.X][pos.Y] != s[0] {
		return nil, false
	}

	// otherwise we are on a matching value, so keep it
	path = append(path, pos)

	// get the remainder of items to check
	r := s[1:]

	// search all around for the rest of the word
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			subPath, ok := g.FindWord(r, Position{pos.X + j, pos.Y + i}, path)

			if ok {
				path = subPath
				result = true
			}
		}
	}

	return path, result
}
