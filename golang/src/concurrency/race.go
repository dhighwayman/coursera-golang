/*
This problem is called a race condition.
Where the output of the program changes based on what line of code executes first.
The order of execution being uncontrollable cause the func calling 'b = a * b'
could happend before or after tham the sentence ' a = b * b'.
*/

package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2

	go func() {
		b = a * b
	}()


	a = b * b
	fmt.Printf("a = %d, b = %d\n", a, b)
}
