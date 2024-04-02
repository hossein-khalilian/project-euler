package main

import (
    "fmt"
    "math/big"
)

func get_first_fib(digits int) int{
    fib1 := big.NewInt(1)
    fib2 := big.NewInt(1)
    ten := big.NewInt(10)
    up_limit := big.NewInt(1)
    for _ = range(digits - 1){
        up_limit = up_limit.Mul(up_limit, ten)
    }

    index := 1
    for {
        index += 1
        fib2 = fib2.Add(fib2, fib1)
        fib1 = fib1.Sub(fib2, fib1)
        if fib2.Cmp(up_limit) > 0{
            break
        }
    }

    return index + 1
}

func main(){
    fmt.Println(get_first_fib(1000))
}