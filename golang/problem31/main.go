package main

import (
    "fmt"
    "strings"
    "strconv"
)

func create_coins_map() map[int]string{
    coins := []int {2, 5, 10, 20, 50, 100, 200}
    coins_map := make(map[int]string, 0)
    for _, coin := range coins{
        s := ""
        for _ = range coin{
            s += "1"
        }
        coins_map[coin] = s
    }

    return coins_map

}

func count_states(s string, coins []int, coins_map map[int]string) int{
    count := 1

    for index, coin := range coins{
        s1 := strings.Replace(s, coins_map[coin], strconv.Itoa(coin), 1)
        if s1 != s{
            // fmt.Println(s1)
            count += count_states(s1, coins[index:], coins_map)
        }
    }


    return count
}

func get_len_states(n int) int{
    s := ""
    for _ = range n{
        s += "1"
    }
    coins := []int{200, 100, 50, 20, 10, 5, 2}
    coins_map := create_coins_map()

    return count_states(s, coins, coins_map)

}


func main(){
    fmt.Println(get_len_states(200))
}