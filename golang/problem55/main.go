package main

import (
	"fmt"
	"math/big"
)

func reverse_int(n big.Int) *big.Int {
	zero := big.NewInt(0)
	one := big.NewInt(1)
	ten := big.NewInt(10)
	n_reverse := big.NewInt(0)
	n_copy := new(big.Int).Set(&n)
	for {
		n_reverse = n_reverse.Add(n_reverse.Mul(n_reverse, ten), one.Mod(n_copy, ten))
		n_copy.Div(n_copy, ten)
		if zero.Cmp(n_copy) == 0 {
			break
		}
	}

	return n_reverse
}

func is_palindromic(n big.Int) bool {
	reverse_n := reverse_int(n)
	if n.Cmp(reverse_n) == 0 {
		return true
	}
	return false
}

func is_lychrel(n big.Int) bool {
	s := new(big.Int).Set(&n)
	i := 0
	for {
		i += 1
		s.Add(s, reverse_int(*s))
		if is_palindromic(*s) {
			// fmt.Println(i)
			return false
		}
		if i >= 50 {
			return true
		}
	}
}

func get_lychrel_below(n int) int {
	total := 0
	for i := range n {
		i += 1
		if is_lychrel(*big.NewInt(int64(i))) {
			total += 1
		}
	}
	return total
}

func get_bigint(n big.Int) *big.Int {
	return &n
}

func main() {
	fmt.Println(get_lychrel_below(10000))
}
