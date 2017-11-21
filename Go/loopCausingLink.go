package main

import (
"fmt"
)

type Link struct {
    data string
    next *Link
}

func buildList() *Link{
    var N,a,b int
    fmt.Scanln(&N)
    arr := make([]Link, N)
    var c rune
    for i, _ := range arr {
        arr[i].data = ""
        for true {
            fmt.Scanf("%c",&c)
            if c == '\n' {
                break
            }
            arr[i].data = arr[i].data + string([]rune{c})
        }
    }
    for true {
        n, _ := fmt.Scanf("%d %d",&a,&b) 
        if n == 0 {
            break
        }
        if a == -1 || b == -1 {
            continue
        }
        arr[a].next = &(arr[b])
    }
    if N == 0 {
        return nil
    }
    return &arr[0]
}

func floyd(s *Link) string {
    var t,h *Link
    if s == nil {
        return "empty"
    }
    t = s
    if t.next == nil {
        return "OK"
    }
    if t.next.next == nil {
        return "OK"
    }
    t = t.next
    h = t.next
    for t!=h {
        if t.next == nil||h.next == nil{
            return "OK"
        }
        if h.next.next == nil {
            return "OK"
        }
        t = t.next
        h = h.next.next
    }
    t = s
    for t!=h {
        t = t.next
        h = h.next
    }
    t = t.next
    for t.next!=h {
        t = t.next
    }
    return t.data
}

func main() {
    fmt.Println(floyd(buildList()))
}
