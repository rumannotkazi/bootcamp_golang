package main

import (
    "hello/internal/hello"
    "os"
)

func main(){
    hello.PrintTo2(os.Stdout)
}