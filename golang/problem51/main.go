package main

import (
    "fmt"
    "time"
    "math"
    "strings"
    "strconv"
)

func is_prime(n int) bool{
    root := int(math.Sqrt(float64(n)))
    for i := range root - 1{
        i += 2
        if n % i == 0{
            return false
        }
    }
    if n == 1 || n == 0{
        return false
    }

    return true
}

func str2int(n_str string) int{
    n, err := strconv.Atoi(n_str)
    if err != nil{
        panic(err)
    }

    return n
}

func get_replacements(n int) [][]int{
    natural_digits := make([]string, 0)
    for i := range 10{
        natural_digits = append(natural_digits, strconv.Itoa(i))
    }

    n_str := strconv.Itoa(n)
    digits := make(map[string]struct{}, 0)
    for _, digit := range(n_str){
        digits[string(digit)] = struct{}{}
    }

    replacements := make([][]int, 0)
    for digit, _ := range digits{
        digit_replacements := make([]int, 0)
        for _, natural_digit := range natural_digits{
            if natural_digit == "0" && digit == string(n_str[0]){
                continue
            }
            num_str :=strings.ReplaceAll(n_str, digit, natural_digit)
            digit_replacements = append(digit_replacements, str2int(num_str))
        }
        replacements = append(replacements, digit_replacements)
    }

    return replacements
}

func get_prime_digit_replacements(n int) int{
    checked_replacements := make(map[string]struct{}, 0)
    i := -1
    for {
        i += 1
        for _, replacement := range get_replacements(i){
            replacement_key := strconv.Itoa(replacement[0]) + strconv.Itoa(replacement[1])
            _, ok := checked_replacements[replacement_key]
            if ok {
                continue
            }
            checked_replacements[replacement_key] = struct{}{}

            family_size := 0
            family := make([]int, 0)
            for _, num := range replacement{
                if is_prime(num){
                    family = append(family, num)
                    family_size += 1
                }
            }
            if family_size >= n{
                fmt.Println(family)
                return family[0]
            }
        }
    }

    return 1
}

func main() {
    start := time.Now()
    fmt.Println(get_prime_digit_replacements(8))
    fmt.Println(time.Since(start))
}