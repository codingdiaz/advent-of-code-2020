package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var expenseList []int

	input, err := os.Open("puzzle-input.txt")
	if err != nil {
		log.Fatalf("error reading input: %s", err.Error())
	}

	scanner := bufio.NewScanner(input)
    for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		expenseList = append(expenseList, i)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	find2(expenseList)
	find3(expenseList)
}

func find2(expenseList []int) {
	var done bool
	for _, x := range expenseList {
		if !done {
			for _, y := range expenseList {
				if (x + y) == 2020 {
					fmt.Printf("Answer to Part One: %v + %v = 2020: %v\n", x, y, (x * y))
					done = true
					break
				}
			}
		}
	}
}

func find3(expenseList []int) {
	var done bool
	for _, x := range expenseList {
		if !done {
			for _, y := range expenseList {
				if !done {
					for _, z := range expenseList {
						if (x + y + z) == 2020 {
							fmt.Printf("Answer to Part Two: %v + %v + %v = 2020: %v\n", x, y, z, (x * y * z))
							done = true
							break
						}
					}
				}
			}
		}
	}
}
