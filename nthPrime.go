package main

import (
"fmt"
)

func main() {
    var N int
    fmt.Scan(&N)
    const uLim = 1000000
    sieve := [uLim]bool{}
    primes := [1011]int{}
    cnt := 1
    for i:=2; i < uLim; i++ {
        if !sieve[i] {
            for j:=i; j < uLim; j+=i {
                sieve[j] = true
            }
            primes[cnt] = i
            if cnt==N {
                fmt.Println(i)
                return
            }
            cnt++
        }
    }
}
