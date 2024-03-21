package main

import (
    "fmt"
    "strings"
)


func number2word(n int) string{
    words_map := make(map[int]string, 0)
    words_map[0] = "zero"
    words_map[1] = "one"
    words_map[2] = "two"
    words_map[3] = "three"
    words_map[4] = "four"
    words_map[5] = "five"
    words_map[6] = "six"
    words_map[7] = "seven"
    words_map[8] = "eight"
    words_map[9] = "nine"
    words_map[10] = "ten"
    words_map[11] = "eleven"
    words_map[12] = "twelve"
    words_map[13] = "thirteen"
    words_map[14] = "fourteen"
    words_map[15] = "fifteen"
    words_map[16] = "sixteen"
    words_map[17] = "seventeen"
    words_map[18] = "eighteen"
    words_map[19] = "nineteen"
    words_map[20] = "twenty"
    words_map[30] = "thirty"
    words_map[40] = "forty"
    words_map[50] = "fifty"
    words_map[60] = "sixty"
    words_map[70] = "seventy"
    words_map[80] = "eighty"
    words_map[90] = "ninety"
    words_map[100] = "hundred"
    words_map[1000] = "thousand"
    words_map[1000000] = "million"


    if n > 0 && n < 20{
        return words_map[n]
    } else if n >= 20 && n < 100{
        if n % 10 == 0{
            return words_map[n]
        }
        return words_map[(n/10)*10] + "-" + number2word(n%10)
    } else if n >= 100 && n < 1000{
        if n % 100 == 0{
            return number2word(n/100) + " " + words_map[100]
        }
        return number2word(n/100) + " " + words_map[100] + " and " + number2word(n%100)
    } else if n >= 1000 && n < 1000000{
        if n % 1000 == 0{
            return number2word(n/1000) + " " + words_map[1000]
        }
        return number2word(n/1000) + " " + words_map[1000] + " and " + number2word(n%1000)
    }

    return words_map[0]

}


func count_letters(n int) int{
    length := 0
    for i := range n{
        i += 1
        s := number2word(i)
        s = strings.ReplaceAll(s, " ", "")
        s = strings.ReplaceAll(s, "-", "")
        length += len(s)
        // fmt.Println(s)

    }

    return length
}


func main(){
    // fmt.Println(number2word(1000))
    fmt.Println(count_letters(1000))
}