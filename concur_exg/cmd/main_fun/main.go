package main

import (
	"concur/internal/concur"
	"fmt"
	"net/http"
	//    "sync"
	//	"fmt"
)

func concurrency_intro() {
	//	c := concur.Initialize()
	//	go c.Increment()
	//	c.Decrement()
	//	fmt.Println(c.Val)

	ch := make(chan int, 10)
	go concur.WriteChannels(&ch, 2, 11)

	//    close(ch)
	concur.ReadChannels(&ch, 5)
	fmt.Println("Next!")
	concur.ReadChannelsClose(&ch)
}

func main() {
	//    url := &concur.Urls{ Urls: map[string]bool{} }
	//    wg:=&sync.WaitGroup{}
	//    wg.Add(1)
	//    go concur.Crawl("https://golang.org/", 4, concur.Ffetcher, wg, url)
	//    wg.Wait()
	ch := make(chan int, 10)
	http.HandleFunc("/work", func(writer http.ResponseWriter, request *http.Request) {
		select {
		case ch <- 2:
			writer.Write([]byte(concur.TheWork()))
			<-ch
		default:
			//close(ch)
			writer.WriteHeader(http.StatusTooManyRequests)
			writer.Write([]byte("Too many reqs"))

		}
	})
	http.ListenAndServe(":8080", nil)
}
