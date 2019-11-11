package main

import "fmt"

func main(){
    var a int
    var b [200]int
    fmt.Scanf("%d", &a)
    for i := 0; i < a; i++ {
        fmt.Scanf("%d", &b[i])
    }
    ans, flag := 0, true
    for flag {
        for i:=0; i<a; i++ {
            if b[i] % 2 > 0 {
                flag = false
            }
            b[i] /= 2
        }
        if flag {
            ans += 1
        }
    }
    fmt.Printf("%d\n", ans)
}
