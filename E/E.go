package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	res_mas := make([]string, t)
	for i := 0; i < t; i++ {
		res_mas[i] = "YES"
		fmt.Fscan(in, &n)
		report_mas := make([]int, n)

		for j := 0; j < n; j++ {
			fmt.Fscan(in, &report_mas[j])
		}
		report_mas_del := make(map[int]int)
		for j := 0; j < n; j++ {
			report_mas_del[report_mas[j]]++
			for j+1 < n && report_mas[j] == report_mas[j+1] {
				j++
			}
		}
		for key := range report_mas_del {
			if report_mas_del[key] > 1 {
				res_mas[i] = "NO"
				break
			}
		}
	}
	for i := 0; i < t; i++ {
		fmt.Println(res_mas[i])
	}
}