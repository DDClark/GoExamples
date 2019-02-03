package main // import "github.com/DDClark/GoExamples/MergeSort"

import "fmt"

func merge(l []int, r []int) []int {
	
}


func mergeSort(input []int) []int {
	l := len(input)
	fmt.Println(input)
	if l == 1 {
		return input
	} else {
		l := mergeSort(input[:l/2])
		r := mergeSort(input[l/2:])
		return merge(l[:],r[:])		
	}
}

func main() {
	fmt.Println("Merge Sort")
	var sample = [...]int {10,3,5,2,9,8,7,1,4,6}
	output := mergeSort(sample[:])
	fmt.Println(sample)
	fmt.Println(output)
}