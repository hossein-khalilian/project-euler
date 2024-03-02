package main

import "fmt"


func sum_even(vector []int) int{
    sum := 0
    for _, v := range vector{
        if v % 2 == 0{
            sum += v            
        }
    }
    return sum
}

func get_fibonachi_sum() int{
    // fib := make([]int, 0)
    // fib = append(fib, 1)
    // fib = append(fib, 1)
    fib := []int{1, 1}
    for {
        fib = append(fib, (fib[len(fib)-1] + fib[len(fib)-2]))
        if fib[len(fib)-1] > 4000000{
            break
        }
    }
    // fmt.Println(fib)
    return sum_even(fib[:len(fib)-1])
}


func get_fibonachi_sum1() int{
    sum := 0
    fib := []int{1, 1}
    for {
        fib[len(fib)-1] = fib[len(fib)-1] + fib[len(fib)-2]
        fib[len(fib)-2] = fib[len(fib)-1] - fib[len(fib)-2]
        if fib[len(fib)-1] > 4000000{
            break
        }
        if fib[len(fib)-1] % 2 == 0{
            sum += fib[len(fib)-1]
        }
    }
    // fmt.Println(fib)
    return sum
}



func main(){
    fmt.Println(get_fibonachi_sum())
    fmt.Println(get_fibonachi_sum1())
}