package main

import (
    "fmt"
    "strconv"
    "strings"
)


func remove_index(vector []int, index int) []int{
    new_vector := make([]int, 0)
    for i, value := range vector{
        if i != index{
            new_vector = append(new_vector, value)
        }
    }

    return new_vector
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

    for index, letter := range letters{
        for _, permutation := range get_permutations(remove_index(letters, index)){
            permutations = append(permutations, letter * pow(10, len(letters)-1) + permutation)
        }
    }

    return permutations
}


func get_pandigitals(n int) []int{
    letters := make([]int, 0)
    for i := range n+1{
        letters = append(letters, i)
    }

    permutations := get_permutations(letters)

    return permutations
}

func get_first_digits(n int) []int{
    n += 1
    letters := make([]int, 0)
    for i := range n{
        letters = append(letters, i)
    }

    permutations := make([]int, 0)
    for i := range n{
        for j := range n - i{
            j += i + 1
            for k := range n - j - 1{
                k += j + 1
                permutations = append(permutations, get_permutations([]int {i, j, k})...)
            }
        }
    }

    return permutations
}

func is_unary_digits(n int) bool{
    s := strconv.Itoa(n)
    for i := range 10{
        s1 := strings.ReplaceAll(s, strconv.Itoa(i), "")
        if len(s) - len(s1) > 1{
            return false
        }
        s = s1
    }
    return true
}

// func get_counts_substring_divisable(n int) int{
//     counts := 0
//     first_digits := get_first_digits(n)
//     primes := []int{2, 3, 5, 7, 11, 13, 17}
//     // primes := []int{17, 13, 11, 7, 5, 3, 2}

//     // fmt.Println(primes)
//     // fmt.Println(first_digits)
//     new_first_digits := first_digits

//     // for index, prime := range primes{
//     //     first_digits = new_first_digits
//     //     new_first_digits = make([]int, 0)
//     //     // fmt.Println(prime)
//     //     for _, digits := range first_digits{
//     //         for i := range n + 1{
//     //             number := digits + i*pow(10, index + 3)
//     //             if is_unary_digits(number){
//     //                 if (number%1000)%prime == 0{
//     //                     new_first_digits = append(new_first_digits, number)
//     //                 }
//     //             }
//     //         }
//     //     }
//     //     break
//     // }

//     for _, digits := range first_digits{
//         for index, prime := range primes{
//             fmt.Println(prime)
//             fmt.Println(index)
//             for i := range n + 1{
//                 digits = digits*10 + i
//                 if is_unary_digits(digits){
//                     if (digits%1000)%prime == 0{
//                         fmt.Println(digits)
//                         break                    
//                     }
//                 }
//             }
//         //     for i := range n + 1{
//         //         number := digits * 10 + i
//         //         if is_unary_digits(number){
//         //             if (number%1000)%prime == 0{
//         //                 fmt.Println(number)
//         //             }
//         //         }
//         //     }
//         //     break
//         // }
//             break
//         }
//         break
//     }
//     // fmt.Println(new_first_digits)
//     counts = len(new_first_digits)

//     return counts
// }


func get_sum_substring_divisable(n int) int{
    sum := 0
    primes := []int{17, 13, 11, 7, 5, 3, 2}
    numbers := get_pandigitals(9)
    for i, number := range numbers{
        flag := true
        for _, prime := range primes{
            if (number%1000) % prime == 0{
                number = number / 10
            } else{
                flag = false
                break
            }
        }
        if flag {
            fmt.Println(numbers[i])
            sum += numbers[i]
        }
    }

    return sum
}


func main(){
    // fmt.Println(len(get_pandigitals(9)))
    // fmt.Println((get_first_digits(9)))
    // fmt.Println((is_unary_digits(11)))
    fmt.Println(get_sum_substring_divisable(9))
    // fmt.Println(get_counts_substring_divisable(9))
}