package main

import (
    "fmt"
    "math"
    "time"
)

func sum_proper_divisors(n int) int{
    sum := 0
    root := int(math.Sqrt(float64(n)))
    for i := range root{
        i += 1
        if n % i == 0{
            sum += i
            sum += n/i
        }
    }
    sum -= n
    if root*root == n{
        sum -= root
    }

    return sum
}

func get_list_abandon_numbers(up int) []int{
    vector := make([]int, 0)
    for i := range up{
        i += 1
        if sum_proper_divisors(i) > i{
            vector = append(vector, i)
        }
    }

    return vector
}

func get_list_non_abandon_sum(up int) []int{
    non_abandon_sum := make([]int, 0)
    abandon_sum := make([]int, 0)
    abandon_sum_bool := make([]bool, up)
    abandon_numbers := get_list_abandon_numbers(up)
    for _, abandon1 := range abandon_numbers{
        for _, abandon2 := range abandon_numbers{
            abandon_sum = append(abandon_sum, abandon1 + abandon2)
            if abandon1 + abandon2 > up{
                break
            }
        }
    }
    // fmt.Println(abandon_sum)
    for _, a := range abandon_sum{
        if a < up{
            abandon_sum_bool[a] = true
        }
    }
    for index, flag := range abandon_sum_bool{
        if flag == false{
            non_abandon_sum = append(non_abandon_sum, index)
        }
    }

    return non_abandon_sum
}

func get_sum_non_abandon_sum(up int) int{
    sum := 0
    non_abandon_sum := get_list_non_abandon_sum(up)
    for _, a := range non_abandon_sum{
        sum += a
    }

    return sum
}

func main(){
    start := time.Now()
    fmt.Println(get_sum_non_abandon_sum(28123))
    fmt.Println(time.Since(start))
}