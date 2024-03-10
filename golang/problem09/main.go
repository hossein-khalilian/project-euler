package main

import (
    "fmt"
    "time"
)

func is_pythagorean(a, b, c int) bool{
    if a*a + b*b - c*c == 0{
        return true
    }

    return false
}

func get_product() int{
 c := 0
 for {
     c += 1
     for b := range c{
         for a := range b{
             if is_pythagorean(a, b, c){
                if a + b + c == 1000{
                    return a*b*c
                }
             }
         }
     }
 }

 return 1
}


func main(){
    // fmt.Println(is_pythagorean(3, 4, 5))
    start := time.Now()
    fmt.Println(get_product())
    fmt.Println(time.Since(start))
}