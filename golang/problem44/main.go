package main

import (
    "fmt"
)

func get_pantagonals_difference() int{
    pantagonals := make([]int, 0)
    pantagonals_map := make(map[int]struct{})
    p := 0
    i := 0
    for {
        i += 1
        p += 3 * i - 2
        pantagonals = append(pantagonals, p)
        pantagonals_map[p] = struct{}{}

        length := len(pantagonals)
        for j := range length - 1{
            index := length - j - 2
            if p >= 2 * pantagonals[index]{
                break
            }

            value1 := pantagonals[index]
            value2 := p - pantagonals[index]
            _, ok := pantagonals_map[value2]
            if ok{
                _, okk := pantagonals_map[value1 - value2]
                if okk {
                    fmt.Println(p, value1, value2, value1 - value2)
                    return value1 - value2
                }
            }
        }
    }

    return 1
}

func main(){
    fmt.Println(get_pantagonals_difference())
}
