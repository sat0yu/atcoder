package main

import "fmt"

func main(){
    var a, b, c, x, res int
    fmt.Scanf("%d", &a)
    fmt.Scanf("%d", &b)
    fmt.Scanf("%d", &c)
    fmt.Scanf("%d", &x)
    for  i:=0; i<a+1; i++ {
        for  j:=0; j<b+1; j++ {
            for  k:=0; k<c+1; k++ {
                if 500*i + 100*j + 50*k == x {
                    res += 1
                }
            }
        }
    }
    fmt.Printf("%d\n", res)
}
