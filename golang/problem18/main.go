package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "time"
)

func sring2int(s string) int{
    v, err := strconv.Atoi(s)
    if err != nil {
        panic(err)
    }

    return v
}

func read_triangle(fileName string) [][]int{
    var traingle [][]int
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
        row := strings.Split(line, " ")
        row_int := make([]int, 0)
        for _, item := range row{
            row_int = append(row_int, sring2int(item))
        }
        traingle = append(traingle, row_int)

    }

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        fmt.Println("Error:", err)
        return nil 
    }

    return traingle
}

func max(vector []int) int{
    max := 0
    for _, v := range vector{
        if v > max{
            max = v
        }
    }
    return max
}

func break_triangle(traingle [][]int) ([][]int, [][]int){
    traingle1 := make([][]int, 0)
    traingle2 := make([][]int, 0)
    for index, vector := range traingle{
        traingle1 = append(traingle1, vector[0:index+1])
        traingle2 = append(traingle2, vector[len(vector)-index - 1:])
    }

    return traingle1, traingle2
}

func find_maximum_path(traingle [][]int) int{
    max_path := 0
    if len(traingle) == 2{
        max_path += max(traingle[0]) + max(traingle[1])
        return max_path
    } else {
        traingle1, traingle2 := break_triangle(traingle[1:])
        // fmt.Println(traingle1)
        // fmt.Println(traingle2)
        // fmt.Println(find_maximum_path(traingle1))
        // fmt.Println(find_maximum_path(traingle2))

        max_paths := []int {find_maximum_path(traingle1), find_maximum_path(traingle2)}
        max_path = max(traingle[0]) + max(max_paths)
    }

    return max_path
}

func main(){
    traingle := read_triangle("triangle.txt")
    // fmt.Println(break_triangle(traingle))
    start := time.Now()
    fmt.Println(find_maximum_path(traingle))
    fmt.Println(time.Since(start))
}