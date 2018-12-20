package insertionsort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func TestInsertionSort(t *testing.T) {
	//
	input := generateSlice(1)
	fmt.Println(input)
	output := insertionSort(input)
	fmt.Println(output)
	//
	input = generateSlice(2)
	fmt.Println(input)
	output = insertionSort(input)
	fmt.Println(output)
	//
	input = generateSlice(6)
	fmt.Println(input)
	output = insertionSort(input)
	fmt.Println(output)
	//
	input = generateSlice(16)
	fmt.Println(input)
	output = insertionSort(input)
	fmt.Println(output)
}
