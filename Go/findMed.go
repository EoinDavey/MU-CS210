package main

import (
"fmt"
"sort"
)

func main() {
    var N int
    fmt.Scan(&N)
    arr := make([]int, N)
    for i:=0; i<N; i++ {
       fmt.Scan(&arr[i])
    }
    sort.Ints(arr)
    if len(arr) % 2 == 1{
        fmt.Printf("%.1f\n",float64(arr[len(arr)/2]))
    } else {
        var tot float64 = float64(arr[(len(arr)-1)/2] + arr[len(arr)/2])
        tot/=2
        fmt.Printf("%.1f\n",tot)
    }
}
