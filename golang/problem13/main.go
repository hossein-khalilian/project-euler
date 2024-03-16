package main

import (
    "fmt"
    "bufio"
    "os"
    "math/big"
)


func get_sum(fileName string) *big.Int{
    sum := big.NewInt(0)
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
        n := new(big.Int)
        n, ok := n.SetString(line, 10)
        if !ok {
            fmt.Println("SetString: error")
            return nil
        }
        sum.Add(sum, n)
    }

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        fmt.Println("Error:", err)
        return nil 
    }

    return sum
}


func print_first_digits(n int, number *big.Int){
    devisor := 1
    for _ = range(n){
        devisor *= 10
    }
    for {
        number.Div(number, big.NewInt(10))
        if number.Cmp(big.NewInt(int64(devisor))) == -1{
            break
        }
    }
    fmt.Println(number)
}


func main(){
	sum := get_sum("numbers.txt")
    print_first_digits(10, sum)
}
