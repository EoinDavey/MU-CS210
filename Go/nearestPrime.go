package main

import (
"fmt"
"math"
)

func abs(i int) int {
    return int(math.Abs(float64(i)))
}

func main() {
    const uLim = 10000
    sieve := [uLim]bool{}
    primes := [10000]int{}
    var N int
    fmt.Scan(&N)
    cnt := 0
    for i:=2; i < uLim; i++ {
        if !sieve[i] {
            for j:=i; j < uLim; j+=i {
                sieve[j] = true
            }
            primes[cnt] = i
            cnt++
        }
    }
    ans := 2
    min := abs(N-ans)
    for _, v := range primes {
        if abs(N-v) < min {
            ans = v
            min = abs(N-v)
        }
    }
    fmt.Println(ans)
}
