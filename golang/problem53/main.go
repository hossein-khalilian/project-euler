package main

import (
    "fmt"
)

func factorial(n int) int{
    if n == 0{
        return 1
    }

    result := 1
    for i := range n{
        i += 1
        result *= i
    }

    return result
}

func is_combinatoric_selection(n int, r int) bool{
    if n - r > r{
        r = n - r
    }

    result := 1
    j := 1
    for i := range n - r {
        i = n - i
        result *= i

        for {
            if j > n - r{
                break
            }
            if result % j == 0{
                result /= j
                j += 1
            } else{
                break
            }
        }

    }
    if result > 1000000{
        return true
    }

    return false
}


func get_combinatoric_selections(n int) int{
    count := 0
    for i := range n{
        i += 1
        for j := range i{
            j += 1
            if is_combinatoric_selection(i, j){
                count += (i - 2*j) + 1
                break
            }
        }
    }

    return count
}

func main(){
    // fmt.Println(factorial(0))
    // fmt.Println(is_combinatoric_selection_01(100, 73))
    fmt.Println(get_combinatoric_selections(100))
}