package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./data.txt")

	var elf int
	var cals int
	elves := make(map[string]int)

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		if elf == 0 {
			elf++
		} else if len(scanner.Text()) == 0 {
			// fmt.Printf("elf: %v  cals: %v \n", elf, cals)
			elves[fmt.Sprintf("elf-%v", elf)] = cals
			elf++
			cals = 0
		} else {
			// fmt.Println(scanner.Text())
			i, _ := strconv.Atoi(scanner.Text())
			cals = cals + i
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	sortedElves := make([]int, 0, len(elves))

	for _, v := range elves {
		sortedElves = append(sortedElves, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sortedElves)))
	fmt.Printf("elves: %v\n", sortedElves)
	fmt.Printf("map len: %v\n", len(elves))
	fmt.Printf("max: %v\n", max(elves))
	top3 := sortedElves[0] + sortedElves[1] + sortedElves[2]
	fmt.Printf("top3: %v\n", top3)
}
func max(numbers map[string]int) int {
	var maxNumber int
	for _, v := range numbers {
		maxNumber = v
		break
	}
	for _, v := range numbers {
		if v > maxNumber {
			maxNumber = v
		}
	}
	return maxNumber
}
