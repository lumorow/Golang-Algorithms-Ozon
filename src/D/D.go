package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n, m, k, min int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	res_matrix := make([][][]int, t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &m)
		res_matrix[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			res_matrix[i][j] = make([]int, m)
			for g := 0; g < m; g++ {
				fmt.Fscan(in, &res_matrix[i][j][g])
			}
		}
		fmt.Fscan(in, &k)
		rotate_data := make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscan(in, &rotate_data[j])
		}
		for j := 0; j < len(rotate_data); j++ {
			for g := 0; g < n; g++ {
				min = res_matrix[i][g][rotate_data[j]-1]
				for c := g + 1; c < n; c++ {
					if min > res_matrix[i][c][rotate_data[j]-1] {
						min = res_matrix[i][c][rotate_data[j]-1]
						for u := c; u > g && u > 0; u-- {
							for p := 0; p < m; p++ {
								res_matrix[i][u][p], res_matrix[i][u-1][p] = res_matrix[i][u-1][p], res_matrix[i][u][p]
							}
						}
					}
				}
			}
		}
	}
	fmt.Println()
	for i := 0; i < t; i++ {
		for j := 0; j < len(res_matrix[i]); j++ {
			for g := 0; g < len(res_matrix[i][j]); g++ {
				fmt.Print(res_matrix[i][j][g], " ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}