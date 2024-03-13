package main

import (
	"fmt"
	"math"
	"time"
)

func get_num_devisor(n int) int{
    if n == 1{
     return 1
    }
	root := int(math.Sqrt(float64(n)))
    num_factors := 0
    for i := range root{
        i += 1
        if n % i == 0{
            num_factors += 2
        }
    }
    return num_factors - 1
}

func get_triangle_num() int{
	i := 0
	triangle_num := 0
	for {
		i += 1
		triangle_num += i
		if get_num_devisor(triangle_num) >= 500{
			break
		}
	}
	return triangle_num
}

func main(){
	start := time.Now()
	fmt.Println(get_triangle_num())
	fmt.Println(time.Since(start))
}