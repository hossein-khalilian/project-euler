package main

import (
    "fmt"
)

func get_digits(n int) [] int{
    digits := make([]int, 0)
    for {
        digits = append(digits, n%10)
        n /= 10
        if n == 0{
            break
        }
    }

    return digits
}

func pow(a int, b int) int{
    result := 1
    for _ = range b{
        result *= a
    }

    return result
}

func get_digits_n_power(n int) int{
    i := 1
    sum := 0
    for i = range 1000000{
        digits := get_digits(i)
        sum_powers := 0
        for _, digit := range digits{
            sum_powers += pow(digit, n) 
        }
        if sum_powers == i{
            sum += i
        }
    }

    return sum - 1
}

func main(){
    fmt.Println(get_digits_n_power(5))
}