/*
Todo:
	break down the string and separate into a list of strings

	Iterate through every char on every list element

	also write functions that check other stuff.


*/

package main

import (
	"strings"
)

var final int = 0
var lines []string
var charType = []string{"corner", "topRail", "botRail", "leftCol", "rightCol", "middle"}

// breaks down input and turns it into a
func organize(input string) {
	lines = strings.Split(strings.TrimSpace(input), "\n")
}

func determineType(lines []string, currentLine int, currentChar int) string {
	switch {
	case currentLine == 0 && currentChar == 0: //topLeft Corner
		return charType[0]
	default:
		return charType[5]
	}

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
