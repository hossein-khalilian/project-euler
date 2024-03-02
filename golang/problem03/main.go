package main

import (
    "fmt"
    "math"
)

func get_primes(n int) []int{
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
    // fmt.Println(get_primes(13195))

    prime_factors := make([]int, 0)
    factors := get_primes(600851475143)
    maxprime := 0
    for _, i := range(factors){
        if len(get_primes(i)) == 2{
            if i > maxprime{
                maxprime = i
            }
            prime_factors = append(prime_factors, i)
        }

    }
    fmt.Println(maxprime)
}