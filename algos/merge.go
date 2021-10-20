package algos

import (
	"fmt"
	"math"
)

func Main() {
	size := 1
	fmt.Println("Enter size of array:")
	fmt.Scanln(&size)

	var numbers []int = make([]int, size+1)
	fmt.Printf("size= %v \n", size)

	fmt.Println("Enter numbers in array:")
	for i := 1; i <= size; i++ {
		fmt.Scanln(&numbers[i])
	}

	n := 1
	mergeSort(n, size+1, numbers)
	fmt.Println("Printing...")
	for i := 1; i <= size; i++ {
		fmt.Printf("%v \n", numbers[i])
	}

}

func mergeSort(n int, size int, array []int) []int {
	if size <= n {
		return array
	} else if n < size {
		// fmt.Printf("size: %v: \n", size)
		q := int(math.Floor(float64((n + size) / 2)))
		mergeSort(n, q, array)
		mergeSort(q+1, size, array)
		merge(n, q, size, array)
	}
	return array

}

func merge(n, q, size int, array []int) []int {
	// fmt.Printf("q: %v and size: %v: \n", q, size)
	n1 := q - n + 1
	n2 := size - q
	// fmt.Printf("%v and %v: \n", n1, n2)
	var L []int = make([]int, n1+1+1)
	var R []int = make([]int, n2+1+1)
	fmt.Printf("first %v and %v: \n", n1, n2)
	for i := 1; i < n1+1; i++ {
		L[i] = array[n+i-1]
	}
	fmt.Printf("second %v and %v: \n", n1, n2)
	fmt.Printf("L length: %d \n", len(L))
	fmt.Printf("R length: %d \n", len(R))
	fmt.Printf("array length: %d \n", len(array))
	fmt.Printf("n2+1: %d \n", n2+1)
	for j := 1; j < n2+1; j++ {
		fmt.Printf("j: %d \n", j)
		fmt.Printf("q+j: %d \n", q+j)
		R[j] = array[q+j]
	}
	fmt.Printf("done \n")
	L[n1+1] = 2147483647
	R[n2+1] = 2147483647 //MAX_INT;
	fmt.Printf("third %v and %v: \n", n1, n2)
	i := 1
	j := 1
	fmt.Printf("n: %v\n", n)
	for k := n; k < size; k++ {
		fmt.Printf("k: %v\n", k)
		if L[i] <= R[j] {
			fmt.Printf("if k: %v\n", k)
			array[k] = L[i]
			i = i + 1
		} else {
			fmt.Printf("else k: %v\n", k)
			array[k] = R[j]
			j = j + 1
		}
	}
	return array
}
