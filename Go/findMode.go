package main

import (
"fmt"
)

func main() {
    var N int
    fmt.Scan(&N)
    arr := make([]int, N)
    for i := 0; i < N; i++ {
        fmt.Scan(&arr[i])
    }
    freqTab := [2002]int{}
    max := 0
    ans := 0
    var v int
    // Reverse traverse list to get values closest to start
    for i:= N-1; i >= 0; i-- {
        v = arr[i]
        freqTab[v+1000]++
        if freqTab[v+1000] >= max {
            ans = v+1000
            max = freqTab[v+1000]
        }
    }
    fmt.Println(ans-1000)
}
