package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"
	"strconv"
)


func main() {

	input, err := os.Open("puzzle-input.txt")
	if err != nil {
		log.Fatalf("error reading input: %s", err.Error())
	}

	scanner := bufio.NewScanner(input)

	var validOne int
	var validTwo int

    for scanner.Scan() {
		d := strings.Split(scanner.Text(), " ")
		
		startToEnd := strings.Split(d[0], "-")

		start, err := strconv.Atoi(startToEnd[0])
		if err != nil {
			log.Fatal(err)
		}

		end, err := strconv.Atoi(startToEnd[1])
		if err != nil {
			log.Fatal(err)
		}

		letter := strings.TrimSuffix(d[1], ":")

		password := d[2]

		var count int
		for _, char := range password {
			if string(char) == letter {
				count ++
			}
		}

		if count >= start && count <= end {
			validOne ++
		}

		var startCorrect, endCorrect bool
		if string(password[start - 1]) == letter {
			startCorrect = true
		}
		if string(password[end - 1]) == letter {
			endCorrect = true
		}

		var thisOneValidTwo bool

		if startCorrect {
			// if true and true
			if endCorrect {
				thisOneValidTwo = false
			} else {
				thisOneValidTwo = true
			}
		} else {
			// if false and true 
			if endCorrect {
				thisOneValidTwo = true
			}
		}

		if thisOneValidTwo {
			validTwo ++
		}


    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
	
	fmt.Println(validOne)
	fmt.Println(validTwo)

}


