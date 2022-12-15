package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// A X > 1 > rock loose
	// B Y > 2 > paper draw
	// C Z > 3 > sissors win
	data, _ := os.ReadFile("./data.txt")
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	var totalPoints int
	for scanner.Scan() {
		value := scanner.Text()
		oValue := string(value[0])
		mySuggestedValue := string(value[2])
		switch mySuggestedValue {
		case "X":
			{ // need to loose against:
				if oValue == "A" { // rock
					// i should be sissors
					totalPoints = totalPoints + 3
				} else if oValue == "B" { // paper
					// should be rock
					totalPoints = totalPoints + 1
				} else if oValue == "C" { // sissor
					// i should be paper
					totalPoints = totalPoints + 2
				}
			}
		case "Y":
			{ // need to tie
				if oValue == "A" { // rock
					// i should be rock
					totalPoints = totalPoints + 1 + 3
				} else if oValue == "B" { // paper
					// should be paper
					totalPoints = totalPoints + 2 + 3
				} else if oValue == "C" { // sissors
					// i should be sissors
					totalPoints = totalPoints + 3 + 3
				}
			}
		case "Z":
			{ // need to win
				if oValue == "A" { // rock
					// i should be paper
					totalPoints = totalPoints + 2 + 6
				} else if oValue == "B" { // paper
					// should be sissors
					totalPoints = totalPoints + 3 + 6
				} else if oValue == "C" { // sissors
					// i should be rock
					totalPoints = totalPoints + 1 + 6
				}

			}
		}
	}
	fmt.Printf("you won a total points of: %v\n", totalPoints)

}
