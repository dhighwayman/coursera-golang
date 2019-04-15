/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order,
separated by a single space on the line.

Your program will define a name struct which has two fields,
fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct
which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file,
your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs
and print the first and last names found in each struct.

*/

package main

import (
  "bufio"
  "fmt"
  "os"
)

type name struct {
    fname string
    lname string
}

func main() {
  names := make([]name,0)
  var fileName string
  fmt.Print("Please enter file name: ")
  fmt.Scan(&fileName)

  file, err := os.Open(fileName)
  if err != nil {
      fmt.Println(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanWords)
  for i:=0 ; scanner.Scan(); i++ {
    var tempName name
    tempName.fname = scanner.Text()
    scanner.Scan()
    tempName.lname = scanner.Text()
    names = append(names, tempName)
  }

  if err := scanner.Err(); err != nil {
      fmt.Println(err)
  }

  for _, n := range names {
    fmt.Print("First Name: " + n.fname)
    fmt.Println(", Last Name: " + n.lname)
  }


}
