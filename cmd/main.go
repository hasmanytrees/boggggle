package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/hasmanytrees/cnd-go-curriculum/packages/boggle/die"
)

var(
	words = []string{}
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

func main(){

	scanner := bufio.NewScanner(os.Stdin)




	go func(){
		for scanner.Scan(){
			usersWord := strings.ToLower(scanner.Text())
			if usersWord != "" {
				words = append(words, usersWord)
			}
		}
	}()

	time.Sleep(10 * time.Second)

	for _, v := range words{
		fmt.Println(v)
	}

}