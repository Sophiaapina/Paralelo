package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if j, ok := seen[complement]; ok {
			return []int{j, i}
		}

		seen[num] = i
	}

	return []int{}
}

func main() {

	var n int
	fmt.Print("Cuantos numeros quieres ingresar: ")
	fmt.Scanln(&n)

	nums := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Numero %d: ", i)
		fmt.Scanln(&nums[i])
	}

	var target int
	fmt.Print("Ingresa el target: ")
	fmt.Scanln(&target)

	result := twoSum(nums, target)

	if len(result) == 0 {
		fmt.Println("No se encontro solucion")
	} else {
		fmt.Println("Indices encontrados:", result)
	}
}
