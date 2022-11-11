package main

import (
    "randomlkjlskj/internal/hello"
    "os"
)

func main(){
    hello.Print("Hello, Xebia")
    printer:= hello.Printer{Output:os.Stderr}
    printer.Print("Error")
    
}