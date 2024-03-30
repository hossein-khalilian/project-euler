package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    // "slices"
    "sort"
)

func read_names(fileName string) []string{
    var matrix []string
    file, err := os.Open(fileName)
    if err != nil {
        fmt.Println("Error:", err)
        return nil
    }
    defer file.Close()

    // Create a scanner to read the file line by line
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        matrix = strings.Split(line, ",")
    }

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        fmt.Println("Error:", err)
        return nil 
    }

    return matrix
}


func word_worth(s string) int{
    worth := 0
    s = strings.Split(s, "\"")[1]
    for _, c := range s{
        worth += int(c - 64)
    }

    return worth
}

func total_scores(names []string) int{
    total_score := 0
    for index, s := range names{
        index += 1
        total_score += index * word_worth(s)
    }

    return total_score
}

func main(){
    names := read_names("0022_names.txt")
    sort.Strings(names)
    fmt.Println(total_scores(names))
}