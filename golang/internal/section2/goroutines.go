package section2

import (
	"fmt"
	"sync"
)

func RunGorountineExercise() {
	array := generateArray(1000000)
	numGoroutines := 4
	sumCh := make(chan int)
	var wg sync.WaitGroup
	chunkSize := len(array) / numGoroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		startIndex := i * chunkSize
		endIndex := (i + 1) * chunkSize
		if i == numGoroutines-1 {
			endIndex = len(array)
		}
		go calculatePartialSum(array, startIndex, endIndex, &wg, sumCh)
	}
	go func() {
		wg.Wait()
		close(sumCh)
	}()
	totalSum := calculateTotalSum(sumCh)
	fmt.Println("Total Sum:", totalSum)
}
func calculateTotalSum(sumCh <-chan int) int {
	totalSum := 0
	for partialSum := range sumCh {
		totalSum += partialSum
	}
	return totalSum
}

func generateArray(size int) []int {
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = i + 1
	}
	return array
}
func calculatePartialSum(array []int, startIndex, endIndex int, wg *sync.WaitGroup, sumCh chan<- int) {
	defer wg.Done()

	partialSum := 0
	for _, value := range array[startIndex:endIndex] {
		partialSum += value
	}

	sumCh <- partialSum
}
