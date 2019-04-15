package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name       string
	food       string
	locomotion string
	sound      string
}

func (cow *Cow) Eat() {
	fmt.Println(cow.food)
}
func (cow *Cow) Move() {
	fmt.Println(cow.locomotion)
}
func (cow *Cow) Speak() {
	fmt.Println(cow.sound)
}

type Bird struct {
	name       string
	food       string
	locomotion string
	sound      string
}

func (bird *Bird) Eat() {
	fmt.Println(bird.food)
}
func (bird *Bird) Move() {
	fmt.Println(bird.locomotion)
}
func (bird *Bird) Speak() {
	fmt.Println(bird.sound)
}

type Snake struct {
	name       string
	food       string
	locomotion string
	sound      string
}

func (snake *Snake) Eat() {
	fmt.Println(snake.food)
}
func (snake *Snake) Move() {
	fmt.Println(snake.locomotion)
}
func (snake *Snake) Speak() {
	fmt.Println(snake.sound)
}

func main() {
	var animals map[string]Animal = make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		request := strings.Split(scanner.Text(), " ")
		if len(request) != 3 {
			fmt.Println("Please enter a command ('newanimal', 'query'), the name of an Animal and the type ('cow', 'bird', 'snake') in case newanimal command or the action ('Eat', 'Move', 'Speak') in case of query command and press ENTER")
			continue
		}
		command := request[0]
		switch command {
		case "newanimal":
			name := request[1]
			animal := animals[name]
			if animal != nil {
				fmt.Println("Animal already exits, please enter other valid name")
				continue
			}
			animaltype := request[2]
			switch animaltype {
			case "cow":
				animals[name] = &Cow{name, "grass", "walk", "moo"}
			case "bird":
				animals[name] = &Bird{name, "worms", "fly", "peep"}
			case "snake":
				animals[name] = &Snake{name, "mice", "slither", "hsss"}
			default:
				fmt.Println("Animal not found, please enter a valid animal ('cow', 'bird', 'snake')")
				continue
			}
			fmt.Println("Created it!")
		case "query":
			name := request[1]
			animal, found := animals[name]
			if !found {
				fmt.Println("Animal not found, please enter a valid name")
				continue
			}
			action := request[2]
			switch action {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("That is not a valid action, please enter a valid action ('Eat', 'Move', 'Speak')")
				continue
			}

		default:
			fmt.Println("enter a valid command ('newanimal', 'query')")
			continue
		}

	}
}
