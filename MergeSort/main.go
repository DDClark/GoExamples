package main // import "github.com/DDClark/GoExamples/MergeSort"

import "fmt"

func merge(l []int, r []int) (result []int) {
	size := len(l) + len(r)
	result = make([]int, size)
	// We need three indices
	i, j, k := 0, 0, 0

	// Both L and R still have elements
	for k < size {
		if i < len(l) && j < len(r) {
			if l[i] < r[j] {
				result[k] = l[i]
				i++
			} else {
				result[k] = r[j]
				j++
			}
			k++
		} else if i < len(l) {
			for ; i < len(l); i++ {
				result[k] = l[i]
				k++
			}
		} else if j < len(r) {
			for ; j < len(r); j++ {
				result[k] = r[j]
				k++
			}
		}
	}
	fmt.Println("Merging")
	fmt.Println(result[:])
	return
}

func mergeSort(input []int) []int {
	size := len(input)
	fmt.Println(input)
	if size == 1 {
		return input
	}

	l := mergeSort(input[:size/2])
	r := mergeSort(input[size/2:])
	return merge(l[:], r[:])

}

func main() {
	fmt.Println("Merge Sort")
	var sample = [...]int{10, 3, 5, 2, 9, 8, 7, 1, 4, 6}
	output := mergeSort(sample[:])
	fmt.Println(sample)
	fmt.Println(output)
}
