package main

import (
	"fmt"
)

func main() {
	var n1, n2,max int32
	fmt.Scanf("%d\n%d", &n1, &n2)
	max = -2500000
	var i, j uint
	for i = 0; i < 32; i++ {
		for j = 0; j < 32; j++ {
			if ((n1 >> i) | (n2 >> j)) > max {
				max = (n1 >> i) | (n2 >> j)
			}
			if ((n1 << i) | (n2 >> j)) > max {
				max = (n1 << i) | (n2 >> j)
			}
			if ((n1 << i) | (n2 << j)) > max {
				max = (n1 << i) | (n2 << j)
			}
			if ((n1 >> i) | (n2 << j)) > max {
				max = (n1 >> i) | (n2 << j)
			}
		}
	}
	fmt.Println(max)
}
