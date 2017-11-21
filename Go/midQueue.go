package main

import (
	"fmt"
)

type Link struct {
	data string
	prev *Link
}

type queue struct {
	front *Link
	back  *Link
    size  int
}

func (q *queue) push(s string) {
	l := new(Link)
    l.data = s
    l.prev = nil
    if q.back == nil {
        q.size = 1
        q.back = l
        q.front = l
    } else {
        q.size++
        q.back.prev = l
        q.back = l
    }
}

func (q *queue) top() string{
    if q.front == nil {
        return ""
    }
    return q.front.data
}

func (q *queue) pop() {
	if q.front == nil {
		return
	}
    q.size--
    if q.front.prev == nil {
        q.front = nil
        q.back = nil
    } else {
        q.front = q.front.prev
    }
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
    if q.size==0 {
        fmt.Println("empty")
    } else {
        for c:=0; c < target; c++ {
            q.pop()
        }
        fmt.Println(q.top())
    }
}
