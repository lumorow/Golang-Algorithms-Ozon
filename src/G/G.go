package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n, m, a, b int
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fmt.Fscan(in, &n, &m)
	friends_map := make(map[int][]int)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &a, &b)
		if _, ok := friends_map[a]; !ok {
			friends_map[a] = make([]int, 0, 5)
		}
		if _, ok := friends_map[b]; !ok {
			friends_map[b] = make([]int, 0, 5)
		}
		friends_map[a] = append(friends_map[a], b)
		friends_map[b] = append(friends_map[b], a)
	}
	for i := 1; i <= n; i++ {
		if _, ok := friends_map[i]; ok {
			var max_count int
			temp_map := make(map[int]int)
			for j := 0; j < len(friends_map[i]); j++ {
				for f := 0; f < len(friends_map[friends_map[i][j]]); f++ {
					temp_map[friends_map[friends_map[i][j]][f]]++
				}
			}
			for j := 0; j < len(friends_map[i]); j++ {
				delete(temp_map, friends_map[i][j])
			}
			delete(temp_map, i)
			if len(temp_map) != 0 {
				for key := range temp_map {
					if max_count < temp_map[key] {
						max_count = temp_map[key]
					}
				}
				temp_keys := make([]int, 0, len(temp_map))
				for key := range temp_map {
					if max_count == temp_map[key] {
						temp_keys = append(temp_keys, key)
					}
				}
				sort.Ints(temp_keys)
				for _, v := range temp_keys {
					fmt.Fprint(out, v, " ")
				}
			} else {
				fmt.Fprint(out, 0)
			}
		} else {
			fmt.Fprint(out, 0)
		}
		fmt.Fprintln(out)
	}
	out.Flush()
}