package main

import (
    "fmt"
    "math"
    "sort"
)

func is_prime(n int) bool{
    root := int(math.Sqrt(float64(n)))
    for i := range root{
        i += 2
        if n % i == 0{
            return false
        }
    }
    if n == 1 {
        return false
    }

    return true
}

func pow(a int, b int) int{
    result := 1
    for _ = range b{
        result *= a
    }

    return result
}

func get_permutations(letters []int) []int{
    permutations := make([]int, 0)
    if len(letters) == 1{
        permutations = append(permutations, letters[0])
        return permutations
    }

    for i, letter := range letters{
        new_letters := make([]int, 0)
        for j := range len(letters){
            if i != j{
                new_letters = append(new_letters, letters[j])
            }
        }
        for _, permutation := range get_permutations(new_letters){
            permutations = append(permutations, letter*pow(10,len(letters)-1) + permutation)
        }
    }

    sort.Sort(sort.Reverse(sort.IntSlice(permutations)))

    return permutations
}

func get_largest_pandigital_prime() int{
    for n := range 10{
        n = 9 - n
        // n += 1
        letters := make([]int, 0)
        for i := range n{
            letters = append(letters, i + 1)
        }
        // fmt.Println(n)
        for _, permutation := range(get_permutations(letters)){
            // fmt.Println(permutation)
            if is_prime(permutation){
                // fmt.Println(permutation)
                return permutation
            }
        }    
    }
    

    return 1
}

func main(){
    fmt.Println(get_largest_pandigital_prime())
}