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
    if n == 1{
        return false
    }

    return true
}

func is_permutation_pair(a int, b int) bool{
    digits_a := make(map[int]int, 0)
    for {
        digits_a[a%10] += 1
        a /= 10
        if a == 0{
            break
        }
    }
    for {
        _, ok := digits_a[b%10]
        if !ok {
            return false
        } else{
            digits_a[b%10] -= 1
        }
        b /= 10
        if b == 0 {
            break
        }
    }
    // fmt.Println(digits_a)
    for _, v := range digits_a{
        if v != 0{
            return false
        }
    }

    return true
}

func get_prime_permutations() int{

    i := 999
    for {
        i += 1
        if i > 10000{
            break
        }

        j := 0
        for {
            j += 2
            if i + 2*j > 10000{
                break
            }
            if i + 3*j < 10000{
                continue
            }

            flag := true
            for k := range 3{
                flag = flag && is_prime(i + k*j)
                if !flag {
                    break
                }
                flag = flag && is_permutation_pair(i, i + k*j)
                if !flag {
                    break
                }
            }
            if flag{
                fmt.Println(i, i + j, i + 2*j)
            }
        }
    }

    return 1
}

func main(){
    get_prime_permutations()
}