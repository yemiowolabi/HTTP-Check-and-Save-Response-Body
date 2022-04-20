////VERSION 2.0.0; Using Channels

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

func checkAndSaveBody(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		//fmt.Println(err)
		s := fmt.Sprintf("%s is DOWN!!!\n", url)
		s += fmt.Sprintf("Error: %v\n", err)
		ch <- s //Sending to channel
	} else {
		defer resp.Body.Close()
		s := fmt.Sprintf("%s ---> status code %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodybytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				//	log.Fatal(err)
				s += "Error Reading Response Body\n"
				ch <- s
			}
			file := strings.Split(url, "//")[1]
			file += ".txt"
			s += fmt.Sprintf("Writing response body to %s\n", file) //To show what it is doing
			err = ioutil.WriteFile(file, bodybytes, 0664)
			if err != nil {
				//log.Fatal(err)
				s += "Error Writing File\n"
				ch <- s
			}
		}
		ch <- s
	}
}

func main() {
	urls := []string{"https://golang.org", "https://oauife.edu.ng", "https://www.googhvgfgle.com/3232.html", "https://medium.com"}
	ch := make(chan string)

	for _, url := range urls {
		go checkAndSaveBody(url, ch)
		fmt.Println(<-ch)

	}
	fmt.Println("Number of GoRoutines:", runtime.NumGoroutine())

}
