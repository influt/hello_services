package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"sync"
	"time"
)

func main(){
	http.HandleFunc("/digits", digitsHandler)
	http.HandleFunc("/chars", charsHandler)
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":8080", nil)
}

func mainHandler(w http.ResponseWriter, _ *http.Request){
	readEndpoint := func(ch chan string, url string) {
                resp, _ := http.Get(url)
		bytes, _ := ioutil.ReadAll(resp.Body)
		ch <- string(bytes)
		close(ch)
	}

	merge := func(channels ...chan string) chan string {
		var wg sync.WaitGroup
		out := make(chan string)
		wg.Add(len(channels))
		for _, ch := range channels {
			go func(ch chan string){
				for char := range ch {
					out <- char
				}
				wg.Done()
			}(ch)

		}
		go func(){
			wg.Wait()
			close(out)
		}()
		return out
	}

	digitCh := make(chan string)
	charCh := make(chan string)

	go readEndpoint(digitCh, "http://127.0.0.1:8080/digits")
	go readEndpoint(charCh, "http://127.0.0.1:8080/chars")

	mergedCh := merge(digitCh, charCh)

	for char := range mergedCh {
		fmt.Fprintf(w, "%s", char)
	}
}

func digitsHandler(w http.ResponseWriter, _ *http.Request){
	for i:=0;i<=9;i++ {
		fmt.Fprintf(w, "%d", i)
	}
	time.Sleep(time.Second * 2)
}

func charsHandler(w http.ResponseWriter, _ *http.Request){
	for i:=65;i<=90;i++ {
		fmt.Fprintf(w, "%c", i)
	}
	time.Sleep(time.Second * 1)
}
