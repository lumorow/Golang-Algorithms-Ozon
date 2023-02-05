package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Data struct {
	input_report string
	left_time    int
	right_time   int
}

func main() {
	var t, n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	res_mas := make([]string, t)
	for i := 0; i < t; i++ {
		res_mas[i] = "YES"
		fmt.Fscan(in, &n)
		report_mas := make([]Data, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &report_mas[j].input_report)
			l1, err := strconv.Atoi(report_mas[j].input_report[0:2])
			if err != nil || l1 > 23 {
				res_mas[i] = "NO"
			}
			r1, err := strconv.Atoi(report_mas[j].input_report[9:11])
			if err != nil || r1 > 23 {
				res_mas[i] = "NO"
			}
			l2, err := strconv.Atoi(report_mas[j].input_report[3:5])
			if err != nil || l2 > 59 {
				res_mas[i] = "NO"
			}
			r2, err := strconv.Atoi(report_mas[j].input_report[12:14])
			if err != nil || r2 > 59 {
				res_mas[i] = "NO"
			}
			l3, err := strconv.Atoi(report_mas[j].input_report[6:8])
			if err != nil || l3 > 59 {
				res_mas[i] = "NO"
			}
			r3, err := strconv.Atoi(report_mas[j].input_report[15:])
			if err != nil || r3 > 59 {
				res_mas[i] = "NO"
			}
			report_mas[j].left_time = l1*3600 + l2*60 + l3
			report_mas[j].right_time = r1*3600 + r2*60 + r3
			if report_mas[j].left_time > report_mas[j].right_time {
				res_mas[i] = "NO"
			}
		}
		if res_mas[i] != "NO" {
			sort.Slice(report_mas, func(i, j int) bool {
				return report_mas[i].left_time < report_mas[j].left_time
			})
			max_time := report_mas[0].right_time
			for j := 0; j < n && res_mas[i] != "NO"; j++ {
				if j+1 < n {
					if report_mas[j+1].left_time <= max_time {
						res_mas[i] = "NO"
						break
					}
					if max_time < report_mas[j+1].right_time {
						max_time = report_mas[j+1].right_time
					}
				}
			}
		}
	}
	for i := 0; i < t; i++ {
		fmt.Println(res_mas[i])
	}
}