package main

import (
"fmt"
"sort"
)

type student struct {
    name string
    score int
}

type ByAge []student

func (a ByAge) Len() int {
    return len(a)
}

func (a ByAge) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
    return a[i].score < a[j].score
}


func main() {
    var N,inn int
    fmt.Scan(&N)
    arr := make([]student, N)
    var ins string
    for i:=0; i < N; i++ {
        fmt.Scan(&ins)
        fmt.Scan(&inn)
        arr[i].name = ins
        arr[i].score = inn
    }
    sort.Sort(ByAge(arr))
    fmt.Println(arr[len(arr)/2].name)
}
