/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

package main

import (
  "bufio"
  "encoding/json"
	"fmt"
  "os"
)

func main() {
  var userMap map[string]string = make(map[string]string)
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print("Please enter your name and press ENTER: ")
  scanner.Scan()
  userMap["name"] = scanner.Text()
  fmt.Print("Please enter your address and press ENTER: ")
  scanner.Scan()
  userMap["address"] = scanner.Text()
  json, err := json.Marshal(userMap)
  if err != nil {
      fmt.Println(err)
  }
  fmt.Println(string(json))
}
