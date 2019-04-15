/*
Write a program to sort an array of integers.
The program should partition the array into 4 parts,
each of which is sorted by a different goroutine.
Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers.
 Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
 When sorting is complete, the main goroutine should print the entire sorted list.
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

	size := len(ints)

	a := ints[(size/4)*0 : (size/4)*1]
	b := ints[(size/4)*1 : (size/4)*2]
	c := ints[(size/4)*2 : (size/4)*3]
	d := ints[(size/4)*3 : (size/4)*4]

	go BubbleSort(a)
	go BubbleSort(b)
	go BubbleSort(c)
	go BubbleSort(d)

	BubbleSort(ints[0:size])
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(ints)
}
