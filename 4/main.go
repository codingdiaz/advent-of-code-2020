package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := os.Open("puzzle-input.txt")
	if err != nil {
		log.Fatalf("error reading input: %s", err.Error())
	}

	passports := getPassports(input)

	var count int
	for _, pass := range passports{
		if validatePassportPartOne(pass) {
			count++
		}
	}

	fmt.Printf("Part One: %v\n", count)

	var count2 int
	for _, pass := range passports{
		if validatePassportPartOne(pass) {
			if validatePassportPartTwo(pass) {
				count2 ++
			}
		}
	}

	fmt.Printf("Part Two: %v\n", count2)

}

func validatePassportPartOne(pass map[string]string) bool {

	// generic validation to complete part 1
	items := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		//"cid",
	}

	for _, i := range items {
		if _, ok := pass[i]; !ok {
			return false
		}
	}

	return true
}

func validatePassportPartTwo(pass map[string]string) bool {

	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	if !checkYears(pass["byr"], 1920, 2002){
		return false
	}

	//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	if !checkYears(pass["iyr"], 2010, 2020){
		return false
	}

	//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	if !checkYears(pass["eyr"], 2020, 2030){
		return false
	}


	//hgt (Height) - a number followed by either cm or in:
	//If cm, the number must be at least 150 and at most 193.
	//If in, the number must be at least 59 and at most 76.
	if !checkHeight(pass["hgt"]) {
		return false
	}


	////hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	if !checkHair(pass["hcl"]) {
		return false
	}

	//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	if !(pass["ecl"] == "amb" || pass["ecl"] == "blu" ||  pass["ecl"] == "brn" || pass["ecl"] == "gry" || pass["ecl"] == "grn" || pass["ecl"] == "hzl" || pass["ecl"] == "oth") {
		return false
	}

	////pid (Passport ID) - a nine-digit number, including leading zeroes.
	if len(pass["pid"]) != 9 {
		return false
	}
	////cid (Country ID) - ignored, missing or not.
	//if pass["cid"][0]

	return true

}

func checkHair(s string) bool{

	////hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	if string(s[0]) != string("#") {
		return false
	}


	if len(s[1:]) != 6 {
		return false
	}

	valid := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}


	for _, x := range s[1:] {
		var v bool 
		for _, y := range valid {
			if string(x) == string(y) {
				v = true
				break
			}
		}

		if v {
			continue 
		} else {
			return false
		}
	}


	return true

}


// checks if a string is 4 characters and between a range of numbers
func checkYears(s string, start, end int) bool {
	if len(s) != 4 {
		return false
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	if i < start || i > end {
		return false
	}

	return true
}

//hgt (Height) - a number followed by either cm or in:
//If cm, the number must be at least 150 and at most 193.
//If in, the number must be at least 59 and at most 76.
func checkHeight(s string) bool {
	last2Characters := s[len(s) - 2:]


	if last2Characters != "cm" && last2Characters != "in" {
		return false
	}

	number, err := strconv.Atoi(s[:len(s) -2])
	if err != nil {
		return false 
	}

	if last2Characters == "cm" {
		if number < 150 || number > 193 {
			return false
		}
	}

	if last2Characters == "in" {
		if number < 59 || number > 76 {
			return false
		}
	}

	return true

}



func getPassports(input io.Reader) []map[string]string {

	var x []string
	var blankLines []int
	var count int

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {

		x = append(x, scanner.Text())

		if scanner.Text() == "" {
			blankLines = append(blankLines, count)
		}

		count ++

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}


	// TODO: this is not very elegant!
	blankLines = append(blankLines, 957)


	var last int
	var passports []map[string]string

	for _, line := range blankLines {

		passport := make(map[string]string)

		for i := last; i < line; i++ {

			var count int
			split := strings.Split(strings.Replace(x[i], " ", ":", -1), ":")
			for count < len(split) {
				passport[split[count]] = split[count+1]
				count = count + 2
			}

		}

		passports = append(passports, passport)

		last = line + 1
	}

	return passports

}
