package main

import (
    "fmt"
    "math"
    "time"
)

func get_factors(n int) []int{
    root := int(math.Sqrt(float64(n)))
    factors := make([]int, 0)
    for i := range root{
        i += 1
        if n % i == 0{
            factors = append(factors, i)
            factors = append(factors, n/i)
        }
    }

    return factors
}

func main(){
    // fmt.Println(get_factors(13195))
    now := time.Now()
    prime_factors := make([]int, 0)
    factors := get_factors(600851475143)
    maxprime := 0
    for _, i := range(factors){
        if len(get_factors(i)) == 2{
            if i > maxprime{
                maxprime = i
            }
            prime_factors = append(prime_factors, i)
        }

    }
    fmt.Println(time.Since(now))
    fmt.Println(maxprime)
}