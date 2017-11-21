package main

import (
"fmt"
"strings"
)

func comp(i,j string) bool {
    if len(i) == len(j) {
        return strings.Compare(i,j) < 0
    }
    return len(i) < len(j)
}

func qSort(arr []string) {
    if len(arr) <= 1 {
        return
    }
    pivot := arr[0]
    store := 1
    for i,v := range arr {
        if comp(v,pivot) {
            arr[i], arr[store] = arr[store], arr[i]
            store++
        }
    }
    arr[0], arr[store-1] = arr[store-1], arr[0]
    qSort(arr[:store-1])
    qSort(arr[store:])
}


func main() {
    var N int
    fmt.Scan(&N)
    arr := make([]string, N)
    for i:=0; i < N; i++ {
        fmt.Scan(&arr[i])
    }
    qSort(arr)
    g := ""
    for _, v := range arr {
        fmt.Print(g, v)
        g=" "
    }
    fmt.Println()
}
