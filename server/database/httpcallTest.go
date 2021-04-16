package main

import (
	"io/ioutil"
	"log"
    "net/http"
    "time"
    "fmt"
    "strconv"
	"sync"
)

func main() {
	testResponseTime()
}

func testResponseTime(){
	var wg  sync.WaitGroup
	wg.Add(1000)
	timebefore := time.Now()
	fmt.Println("Before : ", timebefore.Format(time.ANSIC) + " " + strconv.FormatInt(timebefore.UnixNano(), 10))
	for i := 0; i < 1000; i++ {
		go makeHttpCall(i)
		defer wg.Done()
	}
	wg.Wait()
	timeafter := time.Now()
	fmt.Println("After : ", timeafter.Format(time.ANSIC) + " " + strconv.FormatInt(timeafter.UnixNano(), 10))
}

func makeHttpCall(i int)  (*http.Response, error) {
	return http.Get("http://localhost:8080/headphones")
}

func testResponseContent() {
	resp, err := makeHttpCall(1)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
   	if err != nil {
    log.Fatalln(err)
   	}
	sb := string(body)
   	log.Printf(sb)
}