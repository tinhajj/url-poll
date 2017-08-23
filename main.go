package main

import (
	"bufio"
	"fmt"
	"github.com/andrew-d/go-termutil"
	"net/http"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup //waitgroup for polling results
var results chan string

func main() {
	start := time.Now()

	scanner := bufio.NewScanner(os.Stdin)
	results = make(chan string)

	if !(len(os.Args) > 1) && termutil.Isatty(os.Stdin.Fd()) {
		fmt.Println("No arguments given and nothing to read from stdin")
		return
	}

	if !termutil.Isatty(os.Stdin.Fd()) {
		for scanner.Scan() {
			wg.Add(1)
			go poll(scanner.Text())
		}
	}

	if len(os.Args) > 1 {
		for _, url := range os.Args[1:] {
			wg.Add(1)
			go poll(url)
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for info := range results {
		fmt.Println(info)
	}

	fmt.Println(time.Since(start))
}

func poll(url string) {
	defer wg.Done()
	resp, err := http.Get(url)

	if err != nil {
		results <- err.Error()
		return
	}

	results <- fmt.Sprintf("%-15s%s", resp.Status, url)
}
