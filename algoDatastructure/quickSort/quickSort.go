/*main package implements the quicksort algorithm
for sorting an array
It requires a pivot that the serves as the splitting
point between the array where the sub-arrays, on the left
and right of the pivot can be sorted
*/
package main

import "fmt"

func main() {
	unSortedArray := []int{2, 10, 5, -3, 1, 8, 7, 5, 6}

	arr := quicksort(unSortedArray)

	fmt.Println("sorted Array: ", arr)
}

// quicksort -
func quicksort(arr []int) []int {
	var leftArray []int
	var pivotArray []int
	var rightArray []int

	// base case or index for the recursive call
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < pivot {
			leftArray = append(leftArray, arr[i])

		} else if arr[i] > pivot {
			rightArray = append(rightArray, arr[i])
		} else {
			pivotArray = append(pivotArray, arr[i])
		}

	}
	leftArray = quicksort(leftArray)
	rightArray = quicksort(rightArray)

	// create a slice with the capacity of the size of leftArray
	newSlice := make([]int, len(leftArray))

	// copy everything in leftArray to the new slice
	copy(newSlice, leftArray)

	// spread all the elements in pivot and rightArray to the new slice
	newSlice = append(newSlice, pivotArray...)
	newSlice = append(newSlice, rightArray...)

	return newSlice
}
