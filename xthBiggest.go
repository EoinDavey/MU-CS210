package main

import (
"fmt"
)

func findX(arr []int, L, R, x int) int {
    pivot := arr[L]
    store := L+1
    for i:=store; i < R; i++ {
        if arr[i] > pivot {
            arr[store], arr[i] = arr[i], arr[store]
            store++;
        }
    }
    arr[L], arr[store-1] = arr[store-1], arr[L]
    if store == x {
        return arr[store-1]
    } else if store < x {
        return findX(arr, store, R, x)
    } else {
        return findX(arr, L, store-1, x)
    }
}

func main() {
    var N,x int
    fmt.Scan(&N)
    arr := make([]int, N)
    for i:=0; i < N; i++ {
        fmt.Scan(&arr[i])
    }
    fmt.Scan(&x)
    fmt.Println(findX(arr,0,N,x))
}
