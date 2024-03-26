package main

import (
    "fmt"
    "math/big"
)

func factorial(n int) *big.Int{
    result := big.NewInt(1)
    for i := range n{
        i += 1
        result.Mul(result, big.NewInt(int64(i)))
    }

    return result
}

func sum_digits(x *big.Int) int{
    sum := 0
    ten := big.NewInt(10)
    zero := big.NewInt(0)
    reminder := big.NewInt(1)

    for {
        reminder.Mod(x, ten)
        x.Div(x, ten)
        sum += int(reminder.Int64())
        if x.Cmp(zero) == 0{
            break
        }
    }

    return sum
}

func main(){
    fmt.Println(sum_digits(factorial(100)))
}