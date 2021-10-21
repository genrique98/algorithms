package algos

import (
	"fmt"
)

func GetMinMax() {
	size := 1
	fmt.Println("Enter size of array:")
	fmt.Scanln(&size)

	var numbers []int = make([]int, size)

	fmt.Println("Enter numbers in array:")
	for i := 0; i < size; i++ {
		fmt.Scanln(&numbers[i])
	}

	max_num := numbers[0]
	min_num := numbers[0]

	for i := 1; i < size; i++ {
		if numbers[i] > max_num {
			max_num = numbers[i]
		}
		if numbers[i] < min_num {
			min_num = numbers[i]
		}
	}

	fmt.Printf("max: %v; min: %v \n", max_num, min_num)
}
