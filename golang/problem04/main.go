package main

import (
    "fmt"
    "time"
)

func reverseint(n int) int {
    reverseint := 0
    for {
        reverseint = reverseint * 10 + n % 10
        if n / 10 == 0{
            break
        }
        n = n / 10
    }
    return reverseint
}


func palindromic(digits int) (int, int, int){
    max := 1
    for _ = range digits{
        max *= 10
    }
    max_palind := 0
    factor1 := 0
    factor2 := 0
    for i := range(max){
        for j := range(max){
            i = max - i -1
            j = max - j -1
            mult := i * j
            if mult == reverseint(mult){
                if mult > max_palind{
                    factor1 = i
                    factor2 = j
                    max_palind = mult
                }
            }
        }
        
    }
    return factor1, factor2, max_palind
}

func main(){
    start := time.Now()
    fmt.Println(palindromic(3))
    fmt.Println(time.Since(start))
}