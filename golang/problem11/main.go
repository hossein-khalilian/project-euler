package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)

func sring2int(s string) int{
    v, err := strconv.Atoi(s)
    if err != nil {
        panic(err)
    }

    return v
}

func read_matrix(fileName string) [][]int{
    var matrix [][]int
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
        matrix = append(matrix, row_int)

    }

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        fmt.Println("Error:", err)
        return nil 
    }

    return matrix
}

func get_greatest_product(matrix [][]int) int{
    max_product := 0
    for i := range len(matrix){
        for j := range len(matrix){
            if i + 3 < len(matrix) && j + 3 < len(matrix){
                product := matrix[i][j] * matrix[i+1][j+1] * matrix[i+2][j+2] * matrix[i+3][j+3]
                if product > max_product{
                    max_product = product
                }
            }
            if i - 3 >=0 && j - 3 >= 0{
                product := matrix[i][j] * matrix[i-1][j-1] * matrix[i-2][j-2] * matrix[i-3][j-3]
                if product > max_product{
                    max_product = product
                }
            }
            if i + 3 < len(matrix) && j + 3 < len(matrix) && i - 3 >=0 && j - 3 >= 0{
                product := matrix[i][j] * matrix[i-1][j+1] * matrix[i-2][j+2] * matrix[i-3][j+3]
                if product > max_product{
                    max_product = product
                }
                
                product = matrix[i][j] * matrix[i+1][j-1] * matrix[i+2][j-2] * matrix[i+3][j-3]
                if product > max_product{
                    max_product = product
                }
            }
        }
    }
    return max_product
}

func main(){
    matrix := read_matrix("matrix.txt")
    fmt.Println(get_greatest_product(matrix))
}