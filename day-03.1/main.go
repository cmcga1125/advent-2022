package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./data.txt")
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	var prioritySum int
	lower := "abcdefghijklmnopqrstuvwxyz"
	mixed := lower + strings.ToUpper(lower)

	for scanner.Scan() {
		midpoint := len(scanner.Text()) / 2
		left := scanner.Text()[0:midpoint]
		right := scanner.Text()[midpoint:len(scanner.Text())]
		fmt.Printf("left: %v\n", left)
		fmt.Printf("right: %v\n", right)
		var key rune

		for _, v := range left {
			for _, iv := range right {
				if iv == v {
					key = v
					break
				}
			}
		}
		fmt.Printf("Value: %v\n", string(key))

		value := strings.Index(mixed, string(key)) + 1
		fmt.Printf("value: %v\n", value)
		prioritySum = prioritySum + value
	}
	fmt.Printf("prioritySum: %v\n", prioritySum)
}
