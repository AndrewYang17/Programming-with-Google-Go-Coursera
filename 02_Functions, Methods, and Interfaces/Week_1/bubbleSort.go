/*
Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence
of up to 10 integers. The program should print the integers out on one line, in sorted order,
from least to greatest. Use your favorite search tool to find a description of how the
bubble sort algorithm works.

As part of this program, you should write a function called BubbleSort() which takes a slice
of integers as an argument and returns nothing. The BubbleSort() function should modify the
slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation which swaps the
position of two adjacent elements in the slice. You should write a Swap() function which
performs this operation. Your Swap() function should take two arguments, a slice of
integers and an index value i which indicates a position in the slice.
The Swap() function should return nothing,but it should swap the contents
of the slice in position i with the contents in position i+1.
*/

package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"fmt"
	"strings"
)

func main() {
	var userInput []int
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a sequence of up to 10 integers separated by space: ")
	_ = reader.Scan()

	inputValue := strings.Split(reader.Text(), " ")

	for i := 0; i < len(inputValue); i++ {
		convInputValue, err := strconv.Atoi(inputValue[i])
		if err != nil {
			log.Fatalln("invalid input")
		}
		userInput = append(userInput, convInputValue)
	}

	bubbleSort(userInput)

	for _, v := range userInput {
		fmt.Printf("%d ", v)
	}

	fmt.Println()

}

func bubbleSort(slice []int) {
	sorted := false

	for sorted == false {
		swapped := false
		for i := 0; i < (len(slice) - 1); i++ {
			if slice[i] > slice[i+1] {
				swap(slice, i)
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
	}
}

func swap(slice []int, index int) {
	slice[index], slice[index+1] = slice[index+1], slice[index]
}