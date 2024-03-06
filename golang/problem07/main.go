package main

import (
	"fmt"
	"math"
	"time"
)


func is_prime(n int) bool{
    if len(get_factors(n)) == 2{
        return true
    }
    return false
}

func get_factors(n int) []int{
    root := int(math.Sqrt(float64(n)))
    factors := make([]int, 0)
    if n == 1{
     factors = append(factors, 1)
     return factors
    }
    for i := range root{
        i += 1
        if n % i == 0{
            factors = append(factors, i)
            factors = append(factors, n/i)

        }
    }
    return factors
}

func get_nth_prime(n int) int{
	i := 1
	index := 0
	for {
		i += 1
		if is_prime(i){
			index += 1
		}
		if index == n{
			break
		}
	}

	return i
}

func main(){
	start := time.Now()
	fmt.Println(get_nth_prime(10001))
	fmt.Println(time.Since(start))
}