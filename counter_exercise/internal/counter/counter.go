package counter

import (
    "io"
    "fmt"
    "os"
)

type Counter struct{
    Cnt int
    Output io.Writer
}

func (c *Counter) Next() int{
    temp:=c.Cnt
    c.Cnt++
    return temp
}

func NewCounter() *Counter{
    return &Counter{Cnt:0}
}
func SetNext(usr int) *Counter{
    return &Counter{Cnt:usr}
}
func NewPrinter() *Counter {
    return &Counter{
        Output: os.Stdout,
        }
}
func (c *Counter) Print(val int){
    fmt.Fprintln(c.Output, val)
}

func Print(val int){
    NewPrinter().Print(val)
}

func Run(){

    val:=NewCounter()
    for ;;{
        Print(val.Next())
    }
}