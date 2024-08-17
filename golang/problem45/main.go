package main

import (
    "fmt"
)

func find_triangle_num() int{
    h_max := 0
    p_max := 0
    t_max := 0
    i := 0
    for {
        i += 1
        h_max = i*(2*i-1)

        flag := false
        j := i - 1
        for {
            j += 1
            p_max = j*(3*j-1)/2
            if h_max == p_max{
                flag = true
            }
            if p_max >= h_max{
                break
            }
        }

        if flag{
            for {
                j += 1
                t_max = j*(j+1)/2
                if t_max == h_max{
                    if h_max > 40755{
                        return h_max
                    }
                }
                if t_max >= h_max{
                    break
                }
            }    
        }
    }

    return h_max
}

func main(){
    fmt.Println(find_triangle_num())
}