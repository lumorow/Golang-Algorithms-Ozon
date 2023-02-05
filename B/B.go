
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
	if t >= 1 && t <= 10000 {
		mas := make([]int, t)
		for i := 0; i < t; i++ {
			fmt.Fscan(in, &n)
			mas_temp := make([]int, n)
			for j := 0; j < n; j++ {
				fmt.Fscan(in, &mas_temp[j])
			}
			var map_numbers = make(map[int]int)
			for j := 0; j < n; j++ {
				map_numbers[mas_temp[j]]++
			}
			for key, value := range map_numbers {
				mas[i] += (key * (value - value/3))
			}
		}
		for i := 0; i < t; i++ {
			fmt.Println(mas[i])
		}
	}
}