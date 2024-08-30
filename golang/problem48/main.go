package main

import (
	"fmt"
)

func pow(a int, b int) int{
	result := 1
	for _ = range b{
		result *= a
	}

	return result
}

func customized_pow(a int, b int) int{
	result := 1
	for _ = range b{
		result *= a
		result = result % pow(10,10)
	}

	return result
}

func get_sum_self_powers(n int) int{
	sum := 0
	for i := range n{
		i += 1
		sum += customized_pow(i, i)
		sum = sum % pow(10, 10)
	}

	return sum % pow(10, 10)
}

func main(){
	fmt.Println(get_sum_self_powers(1000))
}