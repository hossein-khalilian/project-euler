package main

import (
    "fmt"
    "math"
)

func unique(vector []int) []int{
    unique_vector := make([]int, 0)
    vector_map := make(map[int]struct{}, 0)
    for _, value := range vector{
        vector_map[value] = struct{}{}
    }
    for k, _ := range vector_map{
        unique_vector = append(unique_vector, k)
    }

    return unique_vector
}

func get_prime_divisors(n int) []int{
    prime_divisors := make([]int, 0)
    root := int(math.Sqrt(float64(n)))
    if n == 1{
        return prime_divisors
    }

    for i := range root{
        i += 2
        if n % i == 0{
            prime_divisors = append(prime_divisors, i)
            prime_divisors = append(prime_divisors, get_prime_divisors(n/i)...)
            break
        }
    }
    if len(prime_divisors) == 0{
        prime_divisors = append(prime_divisors, n)
    }

    return unique(prime_divisors)
}

func get_consecutive_numbers(n int) []int{
    var consecutive_numbers []int

    i := 0
    for {
        consecutive_numbers = nil
        i += 1
        flag := true
        for j := range n{
            consecutive_numbers = append(consecutive_numbers, i+j)
            if len(get_prime_divisors(i+j)) != n{
                flag = false
                break
            }
        }
        if flag{
            break
        }
    }

    return consecutive_numbers
}

func main(){
    fmt.Println(get_consecutive_numbers(4))
}