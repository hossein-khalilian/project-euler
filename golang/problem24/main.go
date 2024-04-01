package main

import (
    "fmt"
)

func get_objects(n int) []int{
    objects := make([]int, 0)
    for i := range n{
        // i += 1
        objects = append(objects, i)
    }

    return objects
}

func pow(n int, p int) int{
    value := 1
    for _ = range p{
        value *= n
    }

    return value
}

func get_permutations(objects []int) []int{
    permutations := make([]int, 0)
    if len(objects) == 1{
        permutations = append(permutations, objects[0])
        return permutations
    } else{
        for _, object := range objects{
            new_objects := make([]int, 0)
            for _, o := range objects{
                if o != object{
                    new_objects = append(new_objects, o)
                }
            }
            for _, permutation := range get_permutations(new_objects){
                permutations = append(permutations, object * pow(10, len(objects)-1) + permutation)
            }
        }
    }

    return permutations
}

func main(){
    objects := get_objects(10)
    fmt.Println(get_permutations(objects)[999999])
}