package main

import (
"fmt"
)

type Link struct {
    data string
    next *Link
    prev *Link
}

func buildList() (*Link, *Link){
    var N,a,b,d int
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
        n, _ := fmt.Scanf("%d %d %d",&a,&b,&d)
        if n == 0 {
            break
        }
        if a == -1 {
            continue
        }
        if b!= -1 {
            arr[a].prev = &(arr[b])
        }
        if d!= -1 {
            arr[a].next = &(arr[d])
        }
    }
    if N == 0 {
        return nil, nil
    }
    return &arr[0], &arr[N-1]
}

func check(s, l *Link) string {
    t := s
    if s == nil {
        return "empty"
    }
    prev := t.prev
    for t.next != nil {
        if t.prev != prev {
            return "FALSE"
        }
        prev = t
        t = t.next
    }
    if t!= l {
        return "FALSE"
    }
    return "TRUE"
}

func main() {
    fmt.Println(check(buildList()))
}
