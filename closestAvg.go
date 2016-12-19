package main

import (
"fmt"
"math"
)

func main() {
    var N int
    fmt.Scan(&N)
    arr := make([]float64,N)
    var tot float64
    for i:=0; i < N; i++ {
        fmt.Scan(&arr[i])
        tot+=arr[i]
    }
    tot/=float64(N)
    ans := arr[0]
    min := math.Abs(tot-ans)
    for _, v := range arr {
        if math.Abs(tot-v) < min {
            ans = v
            min = math.Abs(tot-v)
        }
    }
    fmt.Println(ans)
}

