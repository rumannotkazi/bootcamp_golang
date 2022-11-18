package main

import (
    "counter/internal/counter"
    "fmt"
    "strings"
)


func main(){

    var in_decision string
    iterations:=0
    fmt.Println("Do you want to give input?[y/n]")
    fmt.Scanln(&in_decision)
    if strings.ToLower(in_decision)=="n"{
        val:=counter.NewCounter()
        fmt.Println("How many iterations do you want?")
        fmt.Scanln(&iterations)
        fmt.Println("Starting Counter:...")
        for i:=0;i<iterations;i++{
            counter.Print(val.Next())
        }
    } else if strings.ToLower(in_decision)=="y"{
        fmt.Println("Where do you want the counter to start from?")
        var start_count int
        fmt.Scanln(&start_count)
        val:=counter.SetNext(start_count)
        fmt.Println("How many iterations do you want?")
        fmt.Scanln(&iterations)
        fmt.Println("Starting Counter:...")
        for i:=0;i<iterations;i++{
            //            val.Next()
            counter.Print(val.Next())
        }
    }else{
        counter.Run()
    }





    //
//    val:=counter.NewCounter()
////    counter.Counter{val}
//
//    fmt.Print(val.Next())
//    fmt.Print(val.Next())
//    new_val:=counter.Counter{Cnt:0}
//    new_val:=counter.SetNext(9)
//    fmt.Print(new_val.Next())
//    fmt.Print(new_val.Next())

//    counter:=counter.Counter{cnt:0}
//    fmt.Print(counter.Next(counter))

}