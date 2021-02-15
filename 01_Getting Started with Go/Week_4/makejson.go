//Write a program which prompts the user to first enter a name, and then enter an address.
//Your program should create a map and add the name and address to the map using the
//keys “name” and “address”, respectively. Your program should use Marshal()
//to create a JSON object from the map, and then your program should print the JSON object.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	contact := make(map[string]string)

	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter your name: ")
	_ = reader.Scan()
	contact["name"] = reader.Text()

	fmt.Println("Please enter your address: ")
	_ = reader.Scan()
	contact["address"] = reader.Text()

	result, err := json.Marshal(contact)
	if err != nil {
		log.Fatalln("error in marshaling")
	}

	fmt.Println(string(result))
}
