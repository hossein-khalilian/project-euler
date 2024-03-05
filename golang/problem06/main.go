package main

import (
	"fmt"
)


func square(n int) int{
	return n*n
}


func sum(vector []int) int{
	sum := 0
	for _, v := range vector{
		sum += v
	}
	return sum
}


func sumofsquares(n int) int{
	result := 0
	for i := range n{
		i += 1
		result += square(i)
	}

	return result
}


func squareofsum(n int) int{
	result := 0
	for i := range n{
		i += 1
		result += i
	}
	return square(result)
}


func main(){
	fmt.Println(squareofsum(100) - sumofsquares(100))
}