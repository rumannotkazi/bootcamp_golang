package counter

import (
    "io"
    "fmt"
    "os"
    "time"
)

type Counter struct{
    Cnt int
    Output io.Writer
    Timeout time.Duration
    running bool
}

func (c *Counter) Next() int{
    temp:=c.Cnt
    c.Cnt++
    return temp
}

func NewCounter() *Counter{
    return &Counter{
        Cnt:0,
        Timeout: 3*time.Second,
        running: true,
    }
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

func (c *Counter) runStep(){
    _, err := fmt.Fprintf(c.Output, "%d\n",c.Next())
    if err!=nil{
        panic(err)
    }
    time.Sleep(c.Timeout)
}
func (c *Counter) Run() {
    for c.running{
        c.runStep()
    }
}
func Run(){

//    val:=NewCounter().
//    for ;;{
//        Print(val.Next())
//    }
    NewCounter().Run()
}