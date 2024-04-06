package main

import (
    "fmt"
    "math"
)

func get_divisors(n int) []int{
    divisors := make([]int, 0)
    root := int(math.Sqrt(float64(n)))
    for i := range root{
        i += 1
        if n%i == 0{
            divisors = append(divisors, i)
            if i*i != n{
                divisors = append(divisors, n/i)
            }
        }
    }

    return divisors
}

func is_prime(n int) bool{
    if len(get_divisors(n)) == 2{
        return true
    }
    return false
}

func get_len_quadratic_primes(a int, b int) int{
    lenght := 0
    i := 0
    for {
        i += 1
        if is_prime(i*i + a * i + b){
            lenght += 1
        } else{
            break
        }
    }

    return lenght
}

func get_max_len_quadratic_primes(n int) int{
    max_length := 0
    for i := range 2 * n - 1{
        i -= n - 1
        for j := range 2 * n + 1{
            j -= n
            l := get_len_quadratic_primes(i, j)
            if l > max_length{
                max_length = l
                fmt.Println(i, j, max_length, i*j)
            }
        }
    }

    return max_length
}

func main(){
    // fmt.Println(get_divisors(4))
    // fmt.Println(get_len_quadratic_primes(1, 41))
    fmt.Println(get_max_len_quadratic_primes(1000))
}