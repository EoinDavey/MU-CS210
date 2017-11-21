package main

import (
"fmt"
)

func pal(s string) bool {
    if len(s) <= 1 {
        return true
    }
    if s[0] != s[len(s)-1] {
        return false
    }
    return pal(s[1:len(s)-1])
}

func main() {
    var s string
    fmt.Scan(&s)
    if pal(s) {
        fmt.Println("TRUE")
    } else {
        fmt.Println("FALSE")
    }
}
