/*
Todo:
	break down the string and separate into a list of strings

	Iterate through every char on every list element

	also write functions that check other stuff.


*/

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var counter int = 0 //amt of accessible papers
var lines []string
var charType = []string{"corner", "topLine", "botLine", "leftCol", "rightCol", "middle"}

func main() {
	input, err := os.ReadFile("input") //change this to input when testing the actual case.
	if err != nil {
		log.Fatal(err)
	}
	lines = organize(string(input))

	//for line in lines
	//for char in line
	for currentLine := 0; currentLine < len(lines); currentLine++ {
		for currentChar := 0; currentChar < len(lines[currentLine]); currentChar++ {
			//make if statement to check if current char is roll of paper
			//if yes, then run determine type and determine adjacent, counter also should update
			thisChar := lines[currentLine][currentChar]
			if thisChar == '@' {
				counter = determineAdj(determineType(lines, currentLine, currentChar), counter, lines, currentLine, currentChar)
			}
		}
	}
	fmt.Println("Test")
	fmt.Println(counter)
}

// breaks down input and turns it into a slice
// returns a slice of strings, where each element is a line from the input string.
func organize(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

// determines the positional type of the character being interpreted
// lines = slice of string that is the input in list form
// currentLine = int representing which line we're on in lines[]
// currentChar = int representing which char in currentLine is being analyzed
func determineType(lines []string, currentLine int, currentChar int) string {
	switch {
	case (currentLine == 0 && (currentChar == 0 || currentChar == len(lines[currentLine])-1)) ||
		(currentLine == len(lines)-1 && (currentChar == 0 || currentChar == len(lines[currentLine])-1)): //Should cover corner cases. I know this looks bad
		return charType[0]
	case currentLine == 0 && (currentChar != 0 || currentChar != len(lines[currentLine])): //top line
		return charType[1]
	case currentLine == len(lines)-1 && (currentChar != 0 || currentChar != len(lines[currentLine])): //botLine
		return charType[2]
	case currentChar == 0: //left column
		return charType[3]
	case currentChar == len(lines[currentLine])-1: //right colummn
		return charType[4]
	default: //default is middle
		return charType[5]
	}
}

// determines and counts how many adjacent rolls of paper there are
// counter = int representing the number of retrievable paper objects
// position = type of char being analyzed depending on where it is
// adjacent = number of adjacent paper chars
// lines = slice containing the broken test input
// currentLine = int representing which line we're on in lines[]
// currentChar = int representing which char in currentLine is being analyzed
// returns int = number of retrievable papers, should only increment by 1 every time this is called.
func determineAdj(position string, counter int, lines []string,
	currentLine int, currentChar int) int {
	adjacent := 0 //counts for adjacent rolls of paper
	switch {
	// The below case doesn't need to exist. If it's a corner, it's impossible
	// for there to be >3 adjacent papers.
	//When I comment out this case, the program breaks...
	case position == charType[0]: //corner
		// counter++
	case position == charType[1]: //topline
		adjacent = checkBeforeAndAfter(adjacent, lines, currentLine, currentChar)
		adjacent = checkThreeOnSame(adjacent, lines, currentLine+1, currentChar)
	case position == charType[2]: //bot line
		adjacent = checkThreeOnSame(adjacent, lines, currentLine-1, currentChar)
		adjacent = checkBeforeAndAfter(adjacent, lines, currentLine, currentChar)
	case position == charType[3]: //left column
		adjacent = checkSameAndAfter(adjacent, lines, currentLine-1, currentChar) //iterate i, i+1 for above and below
		adjacent = checkSameAndAfter(adjacent, lines, currentLine+1, currentChar)
		adjacent = checkAfter(adjacent, lines, currentLine, currentChar) //iterate i + 1 on same
	case position == charType[4]: //right column
		adjacent = checkSameAndBefore(adjacent, lines, currentLine-1, currentChar) // checks above line
		adjacent = checkSameAndBefore(adjacent, lines, currentLine+1, currentChar) // checks below line
		adjacent = checkBefore(adjacent, lines, currentLine, currentChar)
	default: //for middle, middle cases
		adjacent = checkThreeOnSame(adjacent, lines, currentLine-1, currentChar)
		adjacent = checkBeforeAndAfter(adjacent, lines, currentLine, currentChar)
		adjacent = checkThreeOnSame(adjacent, lines, currentLine+1, currentChar)
	}

	if adjacent < 4 {
		counter++
	}
	return counter
}

/*
The following functions use these parameters:

adjacent = number of adjacent paper chars
lines = slice containing the broken test input
currentLine = int representing which line we're on in lines[]
currentChar = int representing which char in currentLine is being analyzed

returns int = # of adjacent rolls of paper

*/

func checkThreeOnSame(adjacent int, lines []string, currentLine int, currentChar int) int {
	for i := currentChar - 1; i < currentChar+2; i++ {
		if lines[currentLine][i] == '@' {
			adjacent++
		}
	}
	return adjacent
}

func checkBeforeAndAfter(adjacent int, lines []string, currentLine int, currentChar int) int {
	for i := currentChar - 1; i < currentChar+2; i += 2 {
		if lines[currentLine][i] == '@' {
			adjacent++
		}
	}
	return adjacent
}

func checkSameAndAfter(adjacent int, lines []string, currentLine int, currentChar int) int {
	for i := currentChar; i < currentChar+2; i++ {
		if lines[currentLine][i] == '@' {
			adjacent++
		}
	}
	return adjacent
}

func checkSameAndBefore(adjacent int, lines []string, currentLine int, currentChar int) int {
	for i := currentChar - 1; i < currentChar+1; i++ {
		if lines[currentLine][i] == '@' {
			adjacent++
		}
	}
	return adjacent
}

func checkAfter(adjacent int, lines []string, currentLine int, currentChar int) int {
	if lines[currentLine][currentChar+1] == '@' {
		adjacent++
	}
	return adjacent
}

func checkBefore(adjacent int, lines []string, currentLine int, currentChar int) int {
	if lines[currentLine][currentChar-1] == '@' {
		adjacent++
	}
	return adjacent
}

/*
Identify type of character, if it's paper, do other checks

Identify what kind of piece it is, depending on location

Identify the surrounding pieces




---------

How to parse throught the string

take in the string
	break it into a list of strings
	list[line] //So, every line is an element here

	So I can run string operations to do checks

---


*/
