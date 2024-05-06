package main

import (
	"fmt"
    "bufio"
    "os"
    "math"
    "strings"
)


func read_file(fileName string) []string{
    var words []string
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
        for _, word := range strings.Split(line, ","){
        	words = append(words, strings.Split(word, "\"")[1])
        }

    }

    // Check for errors during scanning
    if err := scanner.Err(); err != nil {
        fmt.Println("Error:", err)
        return nil 
    }

    return words
}

func is_triangle(n int) bool{
	n *= 2
	root := int(math.Sqrt(float64(n)))
	if root * (root+1) == n{
		return true
	}
	return false
}

func get_count_triangle_words() int{
	counts := 0
	words := read_file("0042_words.txt")
	for _, word := range words{
		word_value := 0
		// fmt.Println(word)
		for _, letter := range word{
			word_value += int(letter - 'A' + 1)
		}
		if is_triangle(word_value) {
			counts += 1
		}
	}

	return counts
}

func main(){
	fmt.Println(get_count_triangle_words())
}

