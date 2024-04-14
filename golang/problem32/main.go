package main

import (
    "fmt"
    "strings"
    "strconv"
    "sort"
)

func is_pandigital(s string, n int) bool{
    // fmt.Println(s)
    if len(s) != n{
        return false
    }
    for i := range n{
        i += 1
        s = strings.Replace(s, strconv.Itoa(i), "", 1)
    }
    if len(s) == 0{
        return true
    }

    return false
}

func is_unique_digits(i int, n int) bool{
    s := strconv.Itoa(i)
    for i := range n{
        i += 1
        s1 := strings.Replace(s, strconv.Itoa(i), "", -1)
        if len(s) - len(s1) > 1{
            return false
        }
        s = s1
    }
    if len(s) == 0{
        return true
    }
    return false
}


func get_letters(n int) []int{
    letters := make([]int, 0)
    for i := range n{
        // i = n - i
        i += 1
        letters = append(letters, i)
    }

    return letters
}

func unique(vector []int) []int{
    unique_vector := make([]int, 0)
    unique_vector_map := make(map[int]int, 0)
    for _, v := range vector{
        unique_vector_map[v] = 1
    }
    for unique_value := range unique_vector_map{
        unique_vector = append(unique_vector, unique_value)
    }

    return unique_vector
}

func pow(a int, b int) int{
    result := 1
    for _ = range b{
        result *= a
    }
    return result
}

func RemoveIndex(vector []int, index int) []int {
    new_vector := make([]int, 0)
    for i, v := range vector{
        if i != index{
            new_vector = append(new_vector, v)
        }
    }
    return new_vector
}

func get_objects_permutation(letters []int) []int{
    local_permutations := make([]int, 0)
    if len(letters) == 1{
        local_permutations = append(local_permutations, letters[0])
    } 
    for index, letter := range letters{
        new_letters := RemoveIndex(letters, index)
        for _, permutation := range get_objects_permutation(new_letters){
            local_permutations = append(local_permutations, letter * pow(10, len(letters) - 1) + permutation)
        }
    }

    return local_permutations
}

func get_permutations(n int) []int{
    letters := get_letters(n)

    return get_objects_permutation(letters)
}


func get_unique_digits(n int) []int{
    unique_digits := make([]int, 0)
    permutations := get_permutations(n)
    for _, permutation := range permutations{
        for {
            unique_digits = append(unique_digits, permutation)
            permutation = permutation / 10
            if permutation == 0{
                break
            }
        }
    }

    unique_digits = unique(unique_digits)
    sort.Ints(unique_digits)
    
    return unique_digits
}



func get_sum_pandigital_products(n int) int{
    sum := 0
    products := make([]int, 0)
    unique_digits := make([]int, 0)
    for _, unique_digit := range get_unary_nums(9){
        if unique_digit > 2000{
            break
        }
        unique_digits = append(unique_digits, unique_digit)
    }



    for index, i := range unique_digits{
        if len(strconv.Itoa(i)) >= n{
            break
        }
        for _, j := range unique_digits[index:]{
            s := strconv.Itoa(i) + strconv.Itoa(j)
            if len(s) > n{
                break
            }
            if len(s) < n{
                s = s + strconv.Itoa(i * j)
                if is_pandigital(s, n){
                    fmt.Printf("i = %v, j = %v, product = %v\n", i, j, i*j)
                    products = append(products, i*j)
                    }
            }
        }
    }

    for _, product := range unique(products){
        sum += product
    }

    return sum
}


func get_unary_nums(n int)[]int{
    unary_nums := make([]int, 0)
    letters := get_letters(n)
    for i := range pow(2,n){
        state := strconv.FormatInt(int64(i), 2)
        for _ = range n - len(state){
            state = "0" + state
        }
        new_letters := make([]int, 0)
        for index, s := range state{
            if s == '1'{
                new_letters = append(new_letters, letters[index])
            }
        }
        unary_nums = append(unary_nums, get_objects_permutation(new_letters)...)
    }
    sort.Ints(unary_nums)

    return unary_nums
}


func main(){
    fmt.Println(get_sum_pandigital_products(9))
    // fmt.Println((get_unary_nums(2)))
    // fmt.Println((get_unique_digits(2)))
}