package main

import (
    "fmt"
    "math/big"
)

func get_sum_of_digits(n int) int{
    sum := 0
    value := big.NewInt(1)
    one := big.NewInt(int64(1))
    two := big.NewInt(int64(2))
    ten := big.NewInt(int64(10))
    zero := big.NewInt(int64(0))
    for _ = range n{
        value.Mul(value, two)
    }
    // fmt.Println(value)
    for {
        sum += int(one.Mod(value, ten).Int64())
        value.Div(value, ten)
        if value.Cmp(zero) == 0{
            break
        }
    }


    return sum
}

func main(){
    
    fmt.Println(get_sum_of_digits(1000))
}