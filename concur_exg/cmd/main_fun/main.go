package main

import (
	"concur/internal/concur"
    "fmt"
    "sync"

    //	"fmt"
)

func concurrency_intro(){
    //	c := concur.Initialize()
    //	go c.Increment()
    //	c.Decrement()
    //	fmt.Println(c.Val)

    ch:= make(chan int, 10)
    go concur.WriteChannels(&ch,2,11)

    //    close(ch)
    concur.ReadChannels(&ch,5)
    fmt.Println("Next!")
    concur.ReadChannelsClose(&ch)
}

func main() {
    url := &concur.Urls{ Urls: map[string]bool{} }
    wg:=&sync.WaitGroup{}
    wg.Add(1)
    go concur.Crawl("https://golang.org/", 4, concur.Ffetcher, wg, url)
    wg.Wait()
}


