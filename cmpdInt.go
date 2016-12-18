package main

import (
"fmt"
)

func cmpd(y int, i, x float64) float64 {
    if y == 0 {
        return x
    }
    return (1+(i/100.0)) * cmpd(y-1,i,x)
}

func main() {
    var y int
    var i,x float64
    fmt.Scanf("%d\n%f\n%f",&y,&i,&x)
    fmt.Printf("%.2f\n",cmpd(y,i,x))
}
