/*
Write a program to sort an array of integers. The program should partition the array into 4 parts,
each of which is sorted by a different goroutine. Each partition should be of approximately equal size.
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should
print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	c := make(chan []int)
	s := make([]int, 0)
	userInput := getInput()
	arrays := divideInput(userInput)
	for key := range arrays {
		go sorting(arrays[key], c)
		x := <- c
		s = append(s, x...)
	}
	mergeAndSort(s)
}

func getInput() []int {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Input a series of integers: ")
	_ = reader.Scan()

	userInput := strings.Split(reader.Text(), " ")
	adjUserInput := make([]int, len(userInput))

	for i := range adjUserInput {
		var err error
		adjUserInput[i], err = strconv.Atoi(userInput[i])
		if err != nil {
			log.Fatalln("Input type is not integer")
		}
	}

	if len(adjUserInput) < 4 {
		log.Fatalln("Input length should be equal or more than 4")
	}
	return adjUserInput
}

func sorting(s []int, c chan []int) {
	fmt.Println("Sorting", s, "...")
	sort.Ints(s)
	fmt.Println("Sorted!", s)
	c <- s
}

func divideInput(s []int) map[int][]int {
	arrays := make(map[int][]int)
	counter := 0

	for i := 0; i < len(s); i += 4 {
		arrays[counter] = s[i:checkChunkSize(i+4, len(s))]
		counter++
	}
	return arrays
}

func checkChunkSize(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func mergeAndSort(s []int) {
	sort.Ints(s)
	fmt.Println("The entire sorted list is", s)
}
