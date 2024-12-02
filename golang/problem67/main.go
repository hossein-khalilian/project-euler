package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    // "time"
)

func sring2int(s string) int{
    v, err := strconv.Atoi(s)
    if err != nil {
        panic(err)
    }

    return v
}

func read_triangle(fileName string) [][]int{
    var triangle [][]int
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
        triangle = append(triangle, row_int)

    }

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        fmt.Println("Error:", err)
        return nil 
    }

    return triangle
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

func break_triangle(triangle [][]int) ([][]int, [][]int){
    triangle1 := make([][]int, 0)
    triangle2 := make([][]int, 0)
    for index, vector := range triangle{
        triangle1 = append(triangle1, vector[0:index+1])
        triangle2 = append(triangle2, vector[len(vector)-index - 1:])
    }

    return triangle1, triangle2
}

func find_maximum_path(triangle [][]int, result_chan chan int){
    max_path := 0
    result_chan_local := make(chan int, 2)
    if len(triangle) == 2{
        max_path = max(triangle[0]) + max(triangle[1])
    } else {
        triangle1, triangle2 := break_triangle(triangle[1:])
        go find_maximum_path(triangle1, result_chan_local)
        go find_maximum_path(triangle2, result_chan_local)
        max_path1 := <- result_chan_local
        max_path2 := <- result_chan_local

        max_path = max(triangle[0]) + max([]int {max_path1, max_path2})
    }
    
    result_chan <- max_path
}

func find_maximum_path1(triangle [][]int) int{
    max_path := 0
    if len(triangle) == 2{
        max_path += max(triangle[0]) + max(triangle[1])
        return max_path
    } else {
        triangle1, triangle2 := break_triangle(triangle[1:])
        max_paths := []int {find_maximum_path1(triangle1), find_maximum_path1(triangle2)}
        max_path = max(triangle[0]) + max(max_paths)
    }

    return max_path
}


func find_maximum_path2(triangle [][]int) int{
    max_path := 0
    if len(triangle) == 2{
        max_path += max(triangle[0]) + max(triangle[1])
        return max_path
    } else {
        triangle1, triangle2 := break_triangle(triangle[1:])
        max_paths := []int {find_maximum_path1(triangle1), find_maximum_path1(triangle2)}
        max_path = max(triangle[0]) + max(max_paths)
    }

    return max_path
}

// func break_to_triangles(triangle [][]int) ([][]int, [][]int){
//     triangle1 := make([][]int, 0)
//     triangle2 := make([][]int, 0)
//     for index, vector := range triangle{
//         triangle1 = append(triangle1, vector[0:index+1])
//         triangle2 = append(triangle2, vector[len(vector)-index - 1:])
//     }

//     return triangle1, triangle2
// }


func merge_paths(triangles [][]int) []int{
    result_vector := make([]int, 0)
    for i := range len(triangles)-1{
        for j := range len(triangles[i]){
            triangle := make([][]int, 0)
            triangle = append(triangle, []int {triangles[i][j]})
            triangle = append(triangle, []int {triangles[i+1][j], triangles[i+1][j+1]})
            // fmt.Println(triangle)
            result_vector = append(result_vector, find_maximum_path1(triangle))
        }
    }

    return result_vector
}

func get_max_path(triangle [][]int) [][]int{
    if len(triangle) == 1{
        return triangle
    } else {
        last_row := merge_paths(triangle[len(triangle)-2:])
        triangle = triangle[:len(triangle)-2]
        triangle = append(triangle, last_row)
        // fmt.Println(triangle)
        triangle = get_max_path(triangle)
    }

    
    return triangle
}

func main(){
    triangle := read_triangle("0067_triangle.txt")
    // start := time.Now()

    // fmt.Println(find_maximum_path1(triangle))

    // result_chan := make(chan int, 0)
    // go find_maximum_path(triangle, result_chan)
    // result := <-result_chan
    // fmt.Println(result)

    // fmt.Println(time.Since(start))

    // fmt.Println(triangle[len(triangle)-2:])
    // last_row := merge_paths(triangle[len(triangle)-2:])

    // fmt.Println(triangle)
    // triangle = triangle[:len(triangle)-2]
    // triangle = append(triangle, last_row)
    // fmt.Println(triangle)

    fmt.Println(get_max_path(triangle)[0][0])

}