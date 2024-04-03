package main

import (
    "fmt"
    "math/big"
    "strings"
)

func get_max_length_number(max_n int) int{
    max_len := 0
    max_int := 0
    for i := range max_n{
        i += 1
        length := get_len_recurring_cycle(i)
        // fmt.Println(i, length)
        if length >= max_len{
            max_len = length
            max_int = i
        }
    }

    return max_int
}

func get_len_recurring_cycle(n int) int{
    length := 0
    max_number := big.NewInt(100)
    for {
        max_number = max_number.Mul(max_number, max_number)
        fraction := big.NewInt(1).Div(max_number, big.NewInt(int64(n)))
        length = get_len_recurring(fraction)
        if length >= 0 {
            break
        }
    }

    return length
}

func get_len_recurring(n *big.Int) int{
    divider := big.NewInt(1)
    length := 0
    max_len := -1
    for {
        divider = big.NewInt(1).Mul(divider, big.NewInt(10))
        length += 1
        if big.NewInt(1).Div(n, divider).Cmp(big.NewInt(0)) == 0{
            break
        }

        mod1 := big.NewInt(1).Mod(n, divider)
        mod2 := big.NewInt(1).Mod(big.NewInt(1).Div(n, divider), divider)
        if mod1.Cmp(mod2) == 0{
            len_n := len(n.String())
            len_r := len(strings.ReplaceAll(n.String(), mod1.String(), ""))
            if len_r <= len_n - len_r{
                max_len = length
                if mod1.Int64() == 0{
                    max_len = 0
                }
                break
            }
        }
        
    }

    return max_len
}

func main(){
    // fmt.Println(get_len_recurring_cycle(983))
    fmt.Println(get_max_length_number(1000))
}