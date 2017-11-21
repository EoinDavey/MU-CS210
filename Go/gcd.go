package main

import (
"fmt"
)

func gcd (x, y int) int {
    if y == 0 {
        return x
    }
    return gcd(y, x % y)
}

func main() {
    var x,y int
    fmt.Scanf("%d\n%d",&x,&y)
    fmt.Println(gcd(x,y))
}
