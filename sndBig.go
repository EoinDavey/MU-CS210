package main

import (
"fmt"
)

func main() {
    var N int
    fmt.Scan(&N)
    arr := make([]int, N)
    for i, _ := range arr {
        fmt.Scan(&arr[i])
    }
    mx1 := -250000
    mx2 := -250000
    for _, v := range arr {
        if v >= mx1 {
            mx2 = mx1
            mx1 = v
        } else if v > mx2 {
            mx2 = v
        }
    }
    fmt.Println(mx2)
}
