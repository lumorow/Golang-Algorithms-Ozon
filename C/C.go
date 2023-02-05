package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n, difference_1, difference_2, slice int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	groups := make([][][]int, t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		groups[i] = make([][]int, n/2)
		komanf_numbers := make([]int, n)
		komanf_numbers_i := make([]int, n)
		for t := 0; t < n/2; t++ {
			groups[i][t] = make([]int, 2)
		}
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &komanf_numbers[j])
			komanf_numbers_i[j] = j + 1
		}
		for g := 0; g < n/2; g++ {
			difference_2 = 100
			groups[i][g][0] = komanf_numbers_i[0]
			for k := 1; k < len(komanf_numbers_i); k++ {
				difference_1 = komanf_numbers[0] - komanf_numbers[k]
				if difference_1 < 0 {
					difference_1 *= -1
				}
				if difference_1 == 0 {
					groups[i][g][1] = komanf_numbers_i[k]
					slice = k
					break
				}
				if difference_1 < difference_2 {
					difference_2 = difference_1
					groups[i][g][1] = komanf_numbers_i[k]
					slice = k
				}
			}
			komanf_numbers = append(komanf_numbers[:slice], komanf_numbers[slice+1:]...)
			komanf_numbers = append(komanf_numbers[0:0], komanf_numbers[1:]...)
			komanf_numbers_i = append(komanf_numbers_i[:slice], komanf_numbers_i[slice+1:]...)
			komanf_numbers_i = append(komanf_numbers_i[0:0], komanf_numbers_i[1:]...)
		}
	}
	for i := 0; i < t; i++ {
		for j := 0; j < len(groups[i]); j++ {
			fmt.Println(groups[i][j][0], groups[i][j][1])
		}
		fmt.Println()
	}
}