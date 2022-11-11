package hello

import (
    "fmt"
    "io"
)


func PrintTo2(dest io.Writer){
    fmt.Fprint(dest, "hello, world")
}

