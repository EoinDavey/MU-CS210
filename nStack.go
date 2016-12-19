package main

import (
"fmt"
)

type stack []int

func (s *stack) push(in int) {
    *s = append(*s, in)
}

func (s *stack) pop() {
    if len(*s) > 0 {
        *s = (*s)[:len(*s)-1]
    }
}

func (s *stack) top() int{
    if len(*s) > 0 {
        return (*s)[len(*s)-1]
    } else {
        return 0
    }
}

func main() {
    var N,in int
    var s string
    var st stack
    fmt.Scan(&N)
    for i:=0; i < N; i++ {
        fmt.Scan(&s)
        if s == "PUSH" {
            fmt.Scan(&in)
            if len(st) == 0 {
                st.push(in)
            } else {
                if in > st.top() {
                    st.push(in)
                } else {
                    st.push(st.top())
                }
            }
        } else {
            st.pop()
        }
    }
    if len(st) == 0 {
        fmt.Println("empty")
    } else {
        fmt.Println(st.top())
    }
}
