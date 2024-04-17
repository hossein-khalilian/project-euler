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

func is_prime(n int) bool{
    root := int(math.Sqrt(float64(n)))
    if n == 1{
        return false
    }
    for i := range root - 1{
        i += 2
        if n % i == 0{
            return false
        }
    }

    return true
}

func get_count_circular_primes(n int) int{
    count := 0
    for i := range n{
        i += 1
        if is_prime(i){
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
    }

    return count
}

func main(){
    fmt.Println(get_count_circular_primes(1000000))
}