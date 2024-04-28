package main

import (
    "fmt"
    "strconv"
)

func pow(a int, b int) int{
    result := 1
    for _ = range b{
        result *= a
    }

    return result
}

func get_digits_fraction_product(n int) int{
    product := 1
    index := 1
    for {
        fmt.Println(index)
        fmt.Println(get_digit_fraction(index))
        product *= get_digit_fraction(index)
        index *= 10
        if index > n{
            break
        }
    }

    return product
}

func get_digit_fraction(n int) int{
    index := 0
    step := 0
    steps := make([]int, 0)
    for {
        index += 1
        if n - (9 * pow(10, index-1) * index) < 0{
            break
        }
        step = 9 * pow(10, index-1) * index
        steps = append(steps, step)
    }
    // fmt.Println("n:", n)
    // fmt.Println("index:", index)
    // fmt.Println("step:", step)
    // fmt.Println("steps:", steps)

    for _, step := range steps{
        n = n - step
    }
    i := strconv.Itoa(pow(10, index-1) + (n-1) / index)

    return int(i[(n-1)%index] - '0')
}

func main(){
    // fmt.Println(get_digit_fraction(10000))
    fmt.Println(get_digits_fraction_product(1000000))
}