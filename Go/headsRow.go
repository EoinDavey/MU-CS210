package main

import (
"fmt"
"time"
"math/rand"
)

func main() {
    var row,N,H,tot int
    var reached bool
    fmt.Scanf("%d\n%d",&N,&H)
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i:=0; i < 500000; i++ {
        row = 0
        reached = false
        for j:=0;j < N; j++ {
            roll := r.Float64()
            if roll >= 0.5 {
                row ++
            } else {
                row = 0
            }
            if row == H {
                reached = true
            }
        }
        if row == H {
            reached = true
        }
        if reached {
            tot++
        }
    }
    fmt.Printf("%.0f\n",(float64(tot)/5000.0))
}
