package concur

import (
    "fmt"
    "time"
)

func PrintHello(val int) {

    for i := 0; i < val; i++ {
        fmt.Println("Hello ", i)
        time.Sleep(100 * time.Millisecond)
    }
}

type ConsistentVal struct{
    Val int
}

func Initialize() *ConsistentVal{
    return &ConsistentVal{Val:0}
}

func (c *ConsistentVal) Increment(){
    for i:=0;i<100;i++{
        c.Val++
        time.Sleep(1*time.Nanosecond)
    }
}
func (c *ConsistentVal) Decrement() {
    for i:=0;i<100;i++{
        c.Val--
        time.Sleep(1*time.Nanosecond)
    }
}



// Channels

func WriteChannels(ch *chan int, val int) {
    *ch <- val
}

func ReadChannels(ch *chan int) int{
    d:=<-*ch
    return d
}