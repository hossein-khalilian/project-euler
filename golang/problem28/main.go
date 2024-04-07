package main

import (
	"fmt"
)

func get_diagonals(n int) []int{
	diagonals := make([]int, 0)
	diagonals = append(diagonals, 1)
	step := 0
	for _ = range n * n{
		if len(diagonals) > (n / 2)*4{
			break
		}
		if (len(diagonals) - 1) % 4 == 0{
			step += 2
		}
		diagonals = append(diagonals, diagonals[len(diagonals) - 1] + step)
	}

	return diagonals
}

func get_sum_diagonals(n int) int{
	sum := 0
	diagonals := get_diagonals(n)
	for _, d := range diagonals{
		sum += d
	}

	return sum
}

func main(){
	// fmt.Println(get_diagonals(5))
	fmt.Println(get_sum_diagonals(1001))
}