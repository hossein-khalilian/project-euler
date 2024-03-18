package main

import (
	"fmt"
	"math/big"
)

func combination(start int, end int) *big.Int{
	num := big.NewInt(1)
	div := big.NewInt(1)

	for i := range end-start{
		i += start + 1
		num.Mul(num, big.NewInt(int64(i)))
	}
	
	for i := range(start){
		i += 1
		div.Mul(div, big.NewInt(int64(i)))
	}

	return num.Div(num, div)
}

func num_paths(n int) *big.Int{
	return combination(n, 2*n)
}

func main(){
	fmt.Println(num_paths(20))
}