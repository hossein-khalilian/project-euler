package main

import (
    "fmt"
    "time"
)

func get_len_collatz_sequence(n int, collatz_length map[int]int) int{
    length, ok := collatz_length[n]
    if ok{
        return length
    } else if n == 1{
        length = 0
    } else if n%2 == 0{
        length = get_len_collatz_sequence(n/2, collatz_length)
    } else{
        length = get_len_collatz_sequence(3*n+1, collatz_length)
    }

    collatz_length[n] = 1 + length

    return 1 + length
}

func get_longest_collatz_sequence(max int) (int, int){
    collatz_length := make(map[int]int, 0)
    max_index := 1
    max_length := 1
    for i := range(max){
        i += 1
        length := get_len_collatz_sequence(i, collatz_length)
        if length > max_length{
            max_length = length
            max_index = i
        }
    }
    return max_index, max_length
}

func main(){
    start := time.Now()
    max_index, max_length := get_longest_collatz_sequence(1000000)
    fmt.Println("max_index =", max_index, "\nmax_length =", max_length)
    fmt.Println(time.Since(start))
}