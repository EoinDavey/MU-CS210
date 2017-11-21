package main

import (
"fmt"
)

func div(num,den,n int32) int32 {
    if n == 0 {
        return num/den
    }
    return div((num%den)*10,den,n-1)
}

func main() {
    var num, den, n int32
    fmt.Scanf("%d\n%d\n%d",&num,&den,&n)
    fmt.Println(div(num,den,n))
}
