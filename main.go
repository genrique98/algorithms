package main

import (
	"fmt"

	"github.com/genrique98/algorithms/algos"
)

func main() {
	algorithm := 1
	fmt.Println("Select algorithm:")
	fmt.Println("merge: 1; minmax: 2")
	fmt.Scanln(&algorithm)

	if algorithm == 1 {
		fmt.Println("Starting merge...")
		// setup()
		algos.Main()
	} else if algorithm == 2 {
		fmt.Println("Starting min...")
		algos.Minmax()
	}
}
