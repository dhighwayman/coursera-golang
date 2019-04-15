/*
Write a Bubble Sort program in Go.
The program should prompt the user to type in a sequence of up to 10 integers.
The program should print the integers out on one line, in sorted order, from least to greatest.
Use your favorite search tool to find a description of how the bubble sort algorithm works.

As part of this program, you should write a function called BubbleSort()
which takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation
which swaps the position of two adjacent elements in the slice.
You should write a Swap() function which performs this operation.
Your Swap() function should take two arguments, a slice of integers
and an index value i which indicates a position in the slice.
The Swap() function should return nothing, but it should swap the contents
of the slice in position i with the contents in position i+1.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(ints []int, index int) {
	ints[index], ints[index+1] = ints[index+1], ints[index]
}

func BubbleSort(ints []int) {
	size := len(ints)
	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			if ints[j] > ints[j+1] {
				Swap(ints, j)
			}
		}
	}
}

func main() {
	var ints []int
	var inputSlice []string
	fmt.Print("Enter a sequence of maximum 10 integers: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	inputSlice = strings.Split(input, " ")

	for index, value := range inputSlice {
		if index >= 10 {
			fmt.Println("You have entered more than ten element. Only the first ten will be sorted")
			break
		}

		intValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println(err)
		}
		ints = append(ints, intValue)
	}

	BubbleSort(ints)
	fmt.Println(ints)
}
