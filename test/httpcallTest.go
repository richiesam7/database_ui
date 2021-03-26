package main

import (
	"io/ioutil"
	"log"
    "net/http"
    "time"
    "fmt"
    "strconv"
)

func main() {
	testResponseTime();
}

func testResponseTime() {
	timebefore := time.Now()
	fmt.Println("Before : ", timebefore.Format(time.ANSIC) + " " + strconv.FormatInt(timebefore.UnixNano(), 10))
	for i := 0; i < 1000; i++ {
		http.Get("http://localhost:8080/headphones")
	}
	timeafter := time.Now()
	fmt.Println("After : ", timeafter.Format(time.ANSIC) + " " + strconv.FormatInt(timeafter.UnixNano(), 10))
}

func testResponseContent() {
	resp, err := http.Get("http://localhost:8080/headphones")
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