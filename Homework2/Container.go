package main

import (
	"fmt"
)

func maxAreaWithInfo(height []int) (bestArea int, bestI int, bestJ int) {
	i, j := 0, len(height)-1
	bestArea = 0
	bestI, bestJ = 0, 0

	for i < j {
		h := height[i]
		if height[j] < h {
			h = height[j]
		}

		area := h * (j - i)
		if area > bestArea {
			bestArea = area
			bestI, bestJ = i, j
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return bestArea, bestI, bestJ
}

func main() {
	var n int
	fmt.Print("Cuantos valores tiene height (n): ")
	fmt.Scanln(&n)

	height := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("height[%d]: ", i)
		fmt.Scanln(&height[i])
	}

	bestArea, i, j := maxAreaWithInfo(height)

	// Calculate explanation numbers
	minH := height[i]
	if height[j] < minH {
		minH = height[j]
	}
	width := j - i

	fmt.Println("\n===== RESULTADO =====")
	fmt.Println("Input: height =", height)
	fmt.Println("Output:", bestArea)
	fmt.Printf(
		"Explanation: Se usan las lineas en i=%d (h=%d) y j=%d (h=%d). "+
			"Ancho = %d, altura util = min(%d,%d) = %d, area = %d * %d = %d.\n",
		i, height[i], j, height[j],
		width, height[i], height[j], minH,
		minH, width, bestArea,
	)
}
