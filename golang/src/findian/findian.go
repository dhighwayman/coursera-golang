/*
Write a program which prompts the user to enter a string.

The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.

The program should print “Found!” if the entered string starts with the character ‘i’,
ends with the character ‘n’, and contains the character ‘a’.

The program should print “Not Found!” otherwise.
The program should not be case-sensitive, so it does not matter if the characters
are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings,
“ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var userInput string
  scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Please enter some text and press ENTER: ")
  scanner.Scan()
	userInput = strings.ToLower(scanner.Text())
	var startsWithI bool = strings.HasPrefix(userInput, "i")
	var endsWithN bool = strings.HasSuffix(userInput, "n")
	var containsA bool = strings.Contains(userInput, "a")
	if (startsWithI && endsWithN && containsA) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
