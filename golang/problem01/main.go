package main

import (
    "fmt"
    // "math"
)

func sum(vector []int) int{
    sum := 0
    for _, v := range vector{
        sum += v
    }
    return sum
}

func devidedby35(n int) []int{
    // var vector []int
    vector := make([]int, 0)
    for i := range n{
        if i % 3 == 0 || i % 5 == 0{
            vector= append(vector, i)
        }
    }
        return vector
}

func main(){
    fmt.Println(sum(devidedby35(1000)))
}