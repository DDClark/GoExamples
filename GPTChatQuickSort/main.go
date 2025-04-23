package main

import "fmt"

func quickSort(data []int, c chan []int) {
	if len(data) <= 1 {
		c <- data
		return
	}

	pivot := data[0]
	left := make([]int, 0, len(data))
	right := make([]int, 0, len(data))

	for _, num := range data[1:] {
		if num <= pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	leftCh := make(chan []int)
	rightCh := make(chan []int)
	go quickSort(left, leftCh)
	go quickSort(right, rightCh)

	left = <-leftCh
	right = <-rightCh
	close(leftCh)
	close(rightCh)

	sorted := append(left, pivot)
	sorted = append(sorted, right...)

	c <- sorted
}

func main() {
	data := []int{4, 3, 6, 1, 8, 5, 9, 2}
	c := make(chan []int)
	go quickSort(data, c)
	sortedData := <-c
	fmt.Println(sortedData)
}
