package main

import (
    "fmt"
)

func factorial(n int) int{
    result := 1
    for i := range n{
        i += 1
        result *= i
    }

    return result
}

func get_sum(vector []int) int{
    sum := 0
    for _, v := range vector{
        sum += v
    }

    return sum
}

func get_sum_digit_factorials(n int)int{
    digit_factorials := make([]int, 0)
    for i := range n{
        j := i
        fact := 0
        for {
            fact += factorial(i % 10)
            i = i / 10
            if i == 0{
                break
            }
        }
        if j == fact{
            digit_factorials = append(digit_factorials, j)
        }
        // fmt.Println(j, fact)
    }

    return get_sum(digit_factorials) - 3
}

func main(){
//    fmt.Println(factorial(0))
    fmt.Println(get_sum_digit_factorials(1000000))
}
