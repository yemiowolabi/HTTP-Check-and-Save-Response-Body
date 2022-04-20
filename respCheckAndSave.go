package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is DOWN!!!\n", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s ---> status code %d\n", url, resp.StatusCode)
		if resp.StatusCode == 200 {
			bodybytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			file := strings.Split(url, "//")[1]
			file += ".txt"
			fmt.Printf("Writing response body to %s\n", file) //To show what it is doing
			err = ioutil.WriteFile(file, bodybytes, 0664)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	wg.Done()
}

func main() {
	urls := []string{"https://golang.org", "https://oauife.edu.ng", "https://www.google.com/3232.html", "https://medium.com"}
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
	}
	fmt.Println("Number of GoRoutines:", runtime.NumGoroutine())
	wg.Wait()

}
