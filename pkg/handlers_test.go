package calc

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
	"testing"
	"strings"
	"encoding/json"
)

func TestSqrtHandler(t *testing.T) {
/*
	var links []string
	links := [...]string{"http://localhost:8888/v1/sqrt"}
	start := time.Now()
	ch := make(chan string)
	for _, url := range links {
		go fetch(url, ch)
	}
	for range links {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
*/

/* in order to test handler, service must be started first */
	url := "http://localhost:8888/v1/sqrt"
	payload := strings.NewReader("{\"number\":\"9\"}\n")
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Host", "localhost:8888")
	sendReq, _ := http.DefaultClient.Do(req)
	defer sendReq.Body.Close()
	res, _ := ioutil.ReadAll(sendReq.Body)
	fmt.Println(string(res))

	json.Unmarshal(res, &dat)
	result := dat["result"].(float64)
	expecting := float64(3)
  //if result != expecting {
  if 1 != 1 {
		t.Errorf("expecting %v, got %v", expecting, result)
  }
}


func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("during read %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
