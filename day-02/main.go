package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// A X > 1 > rock
	// B Y > 2 > paper
	// C Z > 3 > sissors
	data, _ := os.ReadFile("./data.txt")
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	var totalPoints int
	for scanner.Scan() {
		value := scanner.Text()
		oValue := string(value[0])
		myValue := string(value[2])
		switch myValue {
		case "X":
			{
				if oValue == "C" {

					totalPoints = totalPoints + 1 + 6
				} else if oValue == "A" {

					totalPoints = totalPoints + 1 + 3
				} else {

					totalPoints = totalPoints + 1
				}
			}
		case "Y":
			{
				if oValue == "A" {
					totalPoints = totalPoints + 2 + 6
				} else if oValue == "B" {

					totalPoints = totalPoints + 2 + 3
				} else {

					totalPoints = totalPoints + 2
				}
			}
		case "Z":
			{
				if oValue == "B" {

					totalPoints = totalPoints + 3 + 6
				} else if oValue == "C" {

					totalPoints = totalPoints + 3 + 3
				} else {

					totalPoints = totalPoints + 3
				}

			}
		}
	}
	fmt.Printf("you won a total points of: %v\n", totalPoints)

}
