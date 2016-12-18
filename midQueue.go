package main

import (
	"fmt"
)

type Link struct {
	data string
	next *Link
	prev *Link
}

type queue struct {
	front *Link
	back  *Link
    size  int
}

func (q *queue) push(s string) {
	l := Link{next: q.back, data: s}
    if q.back == nil {
        q.size = 1
        q.back = &l
        q.front = &l
    } else {
        q.size++
        q.back.prev = &l
        q.back = &l
    }
}

func (q *queue) pop() {
	if q.front == nil {
		return
	}
    q.size--
	q.front = q.front.prev
}

func main() {
    var q queue
    var s string
    for true {
        n, _ := fmt.Scan(&s)
        if n == 0 {
            break
        }
        if s == "INSERT" {
            fmt.Scan(&s)
            q.push(s)
        } else {
            q.pop()
        }
    }
    target := (q.size-1)/2
    t := q.front
    if q.size==0 {
        fmt.Println("empty")
    } else {
        for c:=0; c < target; c++ {
            t = t.prev
        }
        fmt.Println(t.data)
    }
}
