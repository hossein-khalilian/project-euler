package main

import (
    "fmt"
    "math"
)

func is_prime(n int) bool{
    root := int(math.Sqrt(float64(n)))
    for i := range root - 1{
        i += 2
        if n % i == 0{
            return false
        }
    }
    if n == 1{
        return false
    }

    return true
}

func get_primes_lt(n int) []int{
    primes_lt := make([]int, 0)
    if n == 1{
        return primes_lt
    }

    primes_lt = append(primes_lt, 2)
    i := 1
    for {
        i += 2
        if i > n{
            break
        }
        if is_prime(i){
            primes_lt = append(primes_lt, i)
        }
    }

    return primes_lt

}

func get_consecutive_prime_sum(n int) int{
    primes_lt := get_primes_lt(n)
    max_len := 0
    max_prime := 0

    // start := -2
    // // end := 0
    // for {
    //     start += 1
    //     if start >= len(primes_lt){
    //         break
    //     }
    //     num := 0
    //     end := start
    //     for {
    //         end += 1
    //         if end >= len(primes_lt){
    //             break
    //         }
    //         num += primes_lt[end]
    //         if num > n{
    //             break
    //         }
    //         if is_prime(num) && end - start > max_len{
    //             max_len = end - start
    //             max_prime = num
    //             fmt.Println(num, end - start)
    //         }
    //     }
    //     // start += end
    // }

    for start := range len(primes_lt) - 1{
        num := 0
        for i := range len(primes_lt){
            i += start
            num += primes_lt[i]
            if num > n{
                break
            }
            if is_prime(num) && i + 1 - start > max_len{
                max_len = i + 1 - start
                max_prime = num
                fmt.Println(num, i + 1 - start)
            }
        }
    }

    fmt.Println(max_len)

    return max_prime
}


func main(){
    fmt.Println(get_consecutive_prime_sum(1000000))
}