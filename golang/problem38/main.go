package main

import (
    "fmt"
    "strings"
    "strconv"
)

func str2int(s string) int{
    value, err := strconv.Atoi(s)
    if err != nil{
        panic(err)
    }

    return value
}

func is_pandigital(n int, s string) bool{
    for i := range n{
        i += 1
        s1 := strings.ReplaceAll(s, strconv.Itoa(i), "")
        if len(s1) != len(s) - 1{
            return false
        }
        s = s1
    }
    if len(s) != 0{
        return false
    }

    return true
}

func get_largest_pandigital() int{
    largest_pandigital := 0
    for i := range 1000000{
        j := 0
        s := ""
        for{
            j += 1
            s += strconv.Itoa(i * j)
            if len(s) == 9{
                if is_pandigital(9, s){
                    // fmt.Println(i, j)
                    // fmt.Println(s)
                    if largest_pandigital < str2int(s){
                        largest_pandigital = str2int(s)
                    }
                }
            }
            if len(s) > 9{
                break
            }
        }
    }

    return largest_pandigital
}

func main(){
    fmt.Println(get_largest_pandigital())
}