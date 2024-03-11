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


func primes_below(n int) []int{
    primes_lt := make([]int, 0)
    for i := range(n){
        i += 1
        if is_prime(i){
            primes_lt = append(primes_lt, i)
        }
    }
    return primes_lt
}

func sum(vector []int) int{
    sum := 0
    for _, v := range vector{
        sum += v
    }

    return sum
}


func prime_lt(n int) []int{
    root := int(math.Sqrt(float64(n))) + 1
    prime_vector := make([]bool, n+1)
    primes_lt := make([]int, 0)
    prime_vector[0] = true
    prime_vector[1] = true
    for index := range root{
        if prime_vector[index] == false{
            if is_prime(index){
                for i := range(n){
                    i += 2
                    if index*i <= n{
                        prime_vector[index*i] = true
                    } else{
                        break
                    }
                }
            }
        } else{
            continue
        }
    }
    for index, flag := range prime_vector{
        if flag == false{
            primes_lt = append(primes_lt, index)
        }
    }
    // fmt.Println(prime_vector)

    return primes_lt

}


func main(){
    start := time.Now()
    // fmt.Println(sum(primes_below(2000000)))
    fmt.Println(sum(prime_lt(2000000)))
    fmt.Println(time.Since(start))
}