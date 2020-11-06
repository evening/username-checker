package main

import (
	"bufio"
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	t := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		// We use ABSURDLY large keys, and should probably not.
		TLSHandshakeTimeout: 60 * time.Second,
	}
	c := &http.Client{
		Transport: t,
	}

	sem := semaphore.NewWeighted(4)
	url := os.Args[2]
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// var wg sync.WaitGroup
	for scanner.Scan() {
		sem.Acquire(context.Background(), 1)
		// wg.Add(1)
		go func(name string) {
			// defer wg.Done()
			check(url, name, c)
			sem.Release(1)
		}(scanner.Text())
	}
	// wg.Wait()
}

func check(url string, name string, c *http.Client) {
	r, err := c.Get(strings.Replace(url, "{}", name, -1))
	if err != nil {
		log.Panic(err)
		return
	}
	if r.StatusCode == 404 {
		fmt.Printf("%d - %s \n", r.StatusCode, name)

	}
}
