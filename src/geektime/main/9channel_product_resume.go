package main

import (
        "fmt"
        "time"
)

func say(s string) {
        for i := 0; i < 5; i++ {
                time.Sleep(100 * time.Millisecond)
                fmt.Println(s)
        }
}

func sum(s []int, c chan int) {
        sum := 0
        for _, v
}
func main() {
        // go say("world")
        // say("hello")
        s := []int {1,2,3,4,5,6,7,8}
        c = make(chan in)
        go sum(s, )
}