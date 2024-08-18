package main

import (
    "fmt"
    "math"
)

func is_prime(n int) bool{
    root := int(math.Sqrt(float64(n)))
    for i := range root - 1{
        i += 2
        if n % i == 0{
            return false
        }
    }

    return true
}

func is_conjecture_false() int{
    i := 1
    primes := make([]int, 0)
    for{
        i += 2
        if is_prime(i){
            primes = append(primes, i)
        } else{
            flag := false
            for j := range len(primes){
                j = len(primes) - j - 1
                diff := (i - primes[j])/2
                r := int(math.Sqrt(float64(diff)))
                if r*r == diff{
                    flag = true
                    // fmt.Println(i, diff)
                    break
                }
            }
            if !flag{
                // fmt.Println(i)
                return i
            }
        }
    }

    return 1
}

func main(){
    fmt.Println(is_conjecture_false())
}