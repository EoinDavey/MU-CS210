package main

import (
"fmt"
"time"
"math/rand"
)

func main() {
    var cnt,N,H,tot int
    fmt.Scanf("%d\n%d",&N,&H)
    h := float64(H) / 100.0
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for i:=0; i < 500000; i++ {
        cnt = 0
        for j:=0;j < N; j++ {
            roll := r.Float64()
            if roll < h {
                cnt ++
            }
        }
        if float64(cnt) > (float64(N)/2.0) {
            tot++
        }
    }
    fmt.Printf("%.0f\n",(float64(tot)/5000.0))
}
