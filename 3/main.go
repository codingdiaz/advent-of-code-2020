package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	input, err := os.Open("puzzle-input.txt")
	if err != nil {
		log.Fatalf("error reading input: %s", err.Error())
	}

	x := convertToSlice(input)


	fmt.Printf("Part A Total Trees Encountered: %v\n", checkTrees(x, 3, 1))

	fmt.Printf("Total Trees Encountered: %v\n", checkTrees(x, 1, 1))
	fmt.Printf("Total Trees Encountered: %v\n", checkTrees(x, 3, 1))
	fmt.Printf("Total Trees Encountered: %v\n", checkTrees(x, 5, 1))
	fmt.Printf("Total Trees Encountered: %v\n", checkTrees(x, 7, 1))
	fmt.Printf("Total Trees Encountered: %v\n", checkTrees(x, 1, 2))

	fmt.Printf("Part B Answer: %v\n", checkTrees(x, 1, 1) * checkTrees(x, 3, 1) * checkTrees(x, 5, 1) * checkTrees(x, 7, 1) * checkTrees(x, 1, 2))
}

func checkTrees(x [][]string, widthOver, heightDown int) int {

	var treeCounter int

	var width, height int

	for i := 1; i < len(x); i++ {

		width = width + widthOver
		height = height + heightDown

		if height > len(x) {
			break
		}
		
		if x[height][width] == "#" {
			treeCounter++
		}

	}
	return treeCounter
}

func convertToSlice(input io.Reader) [][]string {

	var x [][]string

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "\n")
		var y []string

		for _, char := range line {
			y = append(y, string(char))
		}

		var z []string
		// copy each line a ton so we don't run out of runway
		for i := 1; i < 100; i++ {
			z = append(z, y...)
		}

		x = append(x, z)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return x

}
