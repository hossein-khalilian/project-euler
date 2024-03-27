package main

import (
    "fmt"
    "math"
)

func sum_factors(n int) int{
    if n == 1{
        return 1
    }
    sum := 0
    root := int(math.Sqrt(float64(n)))
    for i := range root{
        i += 1
        if n % i == 0{
         sum += i
         sum += n / i
        }
    }
    if math.Sqrt(float64(n)) == float64(root){
        sum -= root
    }

    return sum - n
}

func get_sum_factors_map(n int) map[int]int{
    factors_map := make(map[int]int)
    for i := range n{
        i += 1
        factors_map[i] = sum_factors(i)
    }

    return factors_map
}

func get_sum_amicable_pairs(factors_map map[int]int) int{
    sum := 0
    for n, s := range factors_map{
        if factors_map[s] == n{
            if n <= s{
                if n != s{
                    // fmt.Println(n, s)
                    sum += n
                    sum += s
                }
            }
        }
    }

    return sum
}

func main(){
    factors_map := get_sum_factors_map(10000)
    // fmt.Println(factors_map)
    fmt.Println(get_sum_amicable_pairs(factors_map))
}