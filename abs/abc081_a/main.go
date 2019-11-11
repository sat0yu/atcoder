package main

import "fmt"

var table = map[string]int{
    "000": 0,
    "001": 1,
    "010": 1,
    "100": 1,
    "101": 2,
    "011": 2,
    "110": 2,
    "111": 3,
}

func main(){
    var a string
    fmt.Scanf("%s", &a)
    fmt.Printf("%d\n", table[a])
}
