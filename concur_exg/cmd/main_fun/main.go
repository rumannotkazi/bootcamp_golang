package main

import (
	"concur/internal/concur"
	"fmt"
)


func main() {
//	c := concur.Initialize()
//	go c.Increment()
//	c.Decrement()
//	fmt.Println(c.Val)

    ch:= make(chan int)
    go concur.WriteChannels(&ch, 8)
    x := concur.ReadChannels(&ch)
    fmt.Println(x)
}
