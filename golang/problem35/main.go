package main

import (
    "fmt"
    "strconv"
    "math"
)

func pow(a int, b int) int{
    result := 1
    for _ = range b{
        result *= a
    }

    return result
}

func get_circulars(n int) []int{
    circulars := make([]int, 0)
    length := len(strconv.Itoa(n))
    multiplier := pow(10, length-1)
    for _ = range length{
        n = (n % 10) * multiplier + n / 10
        circulars = append(circulars, n)
    }
    return circulars
}

func get_divisors(n int) []int{
    root := math.Sqrt(float64(n))
    divisors := make([]int, 0)
    for i := range int(root){
        i += 1
        if n % i == 0{
            divisors = append(divisors, n / i)
            if i * i != n{
                divisors = append(divisors, i)
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

func get_count_circular_primes(n int) int{
 count := 0
 for i := range n{
     circulars := get_circulars(i)
     temp_counter := 0
     for _, circular := range circulars{
        if is_prime(circular) {
            temp_counter += 1
        } else{
            break
        }
     }
     if temp_counter == len(circulars){
        fmt.Println(i)
        count += 1
     }
 }

 return count
}

func main(){
    fmt.Println(get_count_circular_primes(1000000))
}