package main

import (
    "fmt"
)

func get_number_solutions(p int) int{
    number_solutions := 0
    for i := range p{
        i += 1
        for j := range p - i{
            j += i
            k := p - i - j
            if k > j{
                // fmt.Println(i, j, k)
                if i * i + j * j == k * k{
                    // fmt.Println(i,j,k)
                    number_solutions += 1
                }            
            }
        }
    }

    return number_solutions
}

func get_maximized_solution(n int) (int, int){
    maximized_solution := 0
    maximized_p := 0
    for i := range n{
        i += 1
        number_solutions := get_number_solutions(i)
        if maximized_solution < number_solutions{
            // fmt.Println(i)
            maximized_solution  = number_solutions
            maximized_p = i
        }
    }

    return maximized_p, maximized_solution
}

func main(){
    // fmt.Println(get_number_solutions(120))
    fmt.Println(get_maximized_solution(1000))
}