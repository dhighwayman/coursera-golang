/*
Write a program which allows the user to get information about a predefined set of animals.
Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak.
The user can issue a request to find out one of three things about an animal:
1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks.
The following table contains the three animals and their associated data which should be hard-coded into your program.

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program accepts one request at a time from the user, prints out the answer to the request,
and prints out a new prompt. Your program should continue in this loop forever.
Every request from the user must be a single line containing 2 strings.
The first string is the name of an animal, either “cow”, “bird”, or “snake”.
The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each request by printing out the requested data.



*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	sound      string
}

func (animal *Animal) Eat() {
	fmt.Printf("the animal’s food is %s\n", animal.food)
}
func (animal *Animal) Move() {
	fmt.Printf("the animal’s locomotion is %s\n", animal.locomotion)
}
func (animal *Animal) Speak() {
	fmt.Printf("the animal’s sound is %s\n", animal.sound)
}

func main() {
	var animals = map[string]*Animal {
  	"cow": &Animal{"grass", "walk", "moo"},
  	"bird": &Animal{"worms", "fly", "peep"},
  	"snake": &Animal{"mice", "slither", "hsss"},
  }
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		request := strings.Split(scanner.Text(), " ")
    if (len(request) != 2) {
      fmt.Println("Please enter the name of an Animal ('cow', 'bird', 'snake') and an action ('Eat', 'Move', 'Speak') and press ENTER")
      continue
    }
    animalType := request[0]
    action := request[1]
    animal, found := animals[animalType]
    if (!found) {
      fmt.Println("Animal not found, please enter a valid animal ('cow', 'bird', 'snake')")
      continue
    }
		switch action {
		case "Eat":
			animal.Eat()
		case "Move":
			animal.Move()
		case "Speak":
			animal.Speak()
    default:
      fmt.Println("That is not a valid action, please enter a valid action ('Eat', 'Move', 'Speak')")
		}
	}

}
