package main

import (
	"fmt"
	"strconv"
)

func is_palindromes(s string) bool{
	for index := range len(s){
		if s[index] != s[len(s) - 1 - index]{
			return false
		}
	}

	return true
}

func get_sum_palindromes(n int) int{
	sum := 0
	for i := range n{
		if is_palindromes(strconv.Itoa(i)){
			if is_palindromes(strconv.FormatInt(int64(i), 2)){
				// fmt.Println(i)
				sum += i
			}
		}
	}

	return sum
}

func main(){
	// fmt.Println(is_palindromes("5885"))
	fmt.Println(get_sum_palindromes(1000000))
}