package main

import (
    "fmt"
    "math"
)

func get_len_divisors(n int)int{
    root := int(math.Sqrt(float64(n)))
    // divisors := make([]int, 0)
    len_divisors := 0
    for i := range root{
        i += 1
        if n % i == 0{
            len_divisors += 2
            if i * i == n{
                len_divisors -= 1
            }
        }
    }

    return len_divisors
}

func is_prime(n int) bool{
    if get_len_divisors(n) == 2{
        return true
    }
    return false
}

func remove_digits(n int, direction string) []int{
    stages := make([]int, 0)
    if direction == "rtl"{
        for {
            stages = append(stages, n)
            n = n/10    
            if n == 0{
                break
            }
        }
    }
    if direction == "ltr"{
        d := 10
        for {
            stages = append(stages, n % d)
            if n % d == n{
                break
            }
            d *= 10
        }
    }

    return stages
}

func get_stages(n int) []int{
    stages := make([]int, 0)
    stages = append(stages, remove_digits(n, "rtl")...)
    stages = append(stages, remove_digits(n, "ltr")...)

    return stages
}

func get_sum_trancatable_primes()int{
    trancatable_primes := make([]int, 0)
    sum := 0
    i := 10
    for {
        i += 1
        flag := true
        for _, stage := range get_stages(i){
            if !is_prime(stage){
                flag = false
                break
            }
        }
        if flag{
            fmt.Println(i)
            sum += i
            trancatable_primes = append(trancatable_primes, i)
        }
        if len(trancatable_primes) == 11{
            break
        }
    }

    return sum
}

func main(){
    fmt.Println(get_sum_trancatable_primes())
}