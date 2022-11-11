package hello

import (
    "fmt"
    "io"
    "os"
)

type Printer struct {
    Output io.Writer
}

func PrintTo(dest io.Writer){
    fmt.Fprint(dest, "hello, world")
}

func NewPrinter() *Printer {
    return &Printer{
        Output: os.Stdout,
    }
}

func (p *Printer) Print(msg string){
    fmt.Fprint(p.Output, msg)
}

func Print(msg string) {
    NewPrinter().Print(msg)
}