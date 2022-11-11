package hello

import (
    "fmt"
    "os"
)


func PrintTo(dest io.Writer){
    fmt.Fprint(dest, "hello, world")
}

func main(){
    hello.PrintTo(io.Stdout)
}