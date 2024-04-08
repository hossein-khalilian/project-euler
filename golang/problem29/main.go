package main

import (
    "fmt"
    "math/big"
)

func pow(a int, b int) *big.Int{
    result := big.NewInt(1)
    for _ = range b{
        result.Mul(result, big.NewInt(int64(a)))
    }

    return result
}

func get_len_distict_powers(a int, b int) int{
 distict_power_map := make(map[string]struct{})
 for i := range a - 1{
     i += 2
     for j := range b - 1{
         j += 2
         distict_power_map[pow(i,j).String()] = struct{}{}
     }
 }
 // fmt.Println(distict_power_map)

 return len(distict_power_map)
}

func main(){
    fmt.Println(get_len_distict_powers(100, 100))
}