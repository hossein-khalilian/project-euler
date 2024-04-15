package main

import (
    "fmt"
    "strings"
    "strconv"
)

func str2float(s string) float64{
    i, err := strconv.Atoi(s)
    if err != nil {
        panic(err)
    }
    return float64(i)
}

func has_reptitive_runes(a int, b int) bool{
    a1 := strings.Replace(strconv.Itoa(a), strconv.Itoa(a%10), "", 1)
    b1 := strings.Replace(strconv.Itoa(b), strconv.Itoa(a%10), "", 1)
    if str2float(a1) * 10 == float64(a){
        return false
    }
    if (str2float(a1) / str2float(b1)) == float64(a) / float64(b){
        return true
    }

    a1 = strings.Replace(strconv.Itoa(a), strconv.Itoa(a/10), "", 1)
    b1 = strings.Replace(strconv.Itoa(b), strconv.Itoa(a/10), "", 1)
    if (str2float(a1) / str2float(b1)) == float64(a) / float64(b){
        return true
    }

    return false
}

func get_digit_cancelling_fractions() int{
    nums := 1
    denoms := 1
    for i := range 90{
        i += 10
        for j := range 100 - i - 1{
            j += i + 1
            if has_reptitive_runes(i, j){
                nums *= i
                denoms *= j
            }
        }
    }
    // fmt.Println(nums)
    // fmt.Println(denoms)

    return denoms / nums
}

func main(){
    fmt.Println(get_digit_cancelling_fractions())
}