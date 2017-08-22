package main

import (
	"bufio"
	"fmt"
	"github.com/andrew-d/go-termutil"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	scanner := bufio.NewScanner(os.Stdin)

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

	wg.Wait()
	time.Sleep(1000 * time.Millisecond)
}

func poll(url string) {
	defer wg.Done()
	fmt.Println(url)
}
