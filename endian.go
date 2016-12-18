package main

import (
"fmt"
)

func main() {
    var n,filter,b1,b2,b3,b4,res int32
    var s string
    fmt.Scan(&n)
    fmt.Scan(&s)
    filter = 255
    b1 = n & filter
    b2 = (n>>8) & filter
    b3 = (n>>16) & filter
    b4 = (n >> 24) & filter
    var i uint32
    for i = 0; i < 4; i++ {
        switch s[i] {
            case '1':
                res |= b4 << ((3-i) * 8)
            case '2':
                res |= b3 << ((3-i) * 8)
            case '3':
                res |= b2 << ((3-i) * 8)
            case '4':
                res |= b1 << ((3-i) * 8)
        }
    }
    fmt.Println(res)
}
