package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, a, b int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)
	if n >= 1 && n <= 10000 {
		mas := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a, &b)
			if a >= -1000 && a <= 1000 && b >= -1000 && b <= 1000 {
				mas[i] = a + b
			} else {
				i--
				n--
			}
		}
		for i := 0; i < n; i++ {
			fmt.Println(mas[i])
		}
	}
}