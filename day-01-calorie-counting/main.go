package main

import (
	"fmt"
	"log"
	"bufio"
	// "strconv"
	"strings"
)



func main() {
	data :=`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

	var elf int
	var cals int

	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		if elf == 0 {
			elf++
		} else if len(scanner.Text()) == 0 {
			fmt.Printf("elf: %v  cals: %v \n", elf, cals)
			elf++
			cals = 0
		} else {
			fmt.Println(scanner.Text())
			// val := scanner.Text
			// i, _ := strconv.Atoi(string(val))
			// cals = cals + i
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
