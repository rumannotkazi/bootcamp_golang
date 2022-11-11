package hello

import (
    "fmt"
    "io"
)


func PrintTo(dest io.Writer){
    fmt.Fprint(dest, "hello, world")
}

