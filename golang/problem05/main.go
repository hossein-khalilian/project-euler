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


func primes_lt(n int) []int{
    primes_lt := make([]int, 0)
    for i := range(n){
        i += 1
        if len(get_factors(i)) == 2{
            primes_lt = append(primes_lt, i)
        }
    }
    return primes_lt
}


func prime_factors(n int) []int{
    factors := get_factors(n)
    prime_factors := make([]int, 0)
    for _, factor := range factors{
        if is_prime(factor){
            prime_factors = append(prime_factors, factor)
        }
    }

    return prime_factors
}


func prime_factors_degree(n int) map[int]int{
    prime_factors_map := make(map[int]int, 0)
    prime_factors := prime_factors(n)
    for _, prime_factor := range prime_factors{
        for {
            if n % prime_factor == 0{
                prime_factors_map[prime_factor] += 1
                n /= prime_factor
            } else{
                break
            }
        }
    }

    return prime_factors_map
}


func get_evenly_divisable(n int) int{
    max_primes := make(map[int]int, 0)
    for i := range n{
        prime_factors_map := prime_factors_degree(i)
        for prime, degree := range prime_factors_map{
            v, ok := max_primes[prime]
            if !ok{
                max_primes[prime] = degree
            } else if v < degree{
                max_primes[prime] = degree
            }
        }
    }
    value := 1
    for prime, degree := range max_primes{
        for _ = range degree{
            value *= prime
        }
    }


    return value
}


func get_prime_factors(n int, prime_factors map[int]int){
    factors := get_factors(n)
    i := 0
    for _, factor := range factors{
        i += 1
        if is_prime(factor){
            prime_factors[factor] += 1
        } else if factor == n || factor == 1{
            i -= 1
            continue
        } else if i < 3{
            get_prime_factors(factor, prime_factors)
        }
    }
}


func main(){
    start := time.Now()
    fmt.Println(get_evenly_divisable(20))
    fmt.Println(time.Since(start))

}