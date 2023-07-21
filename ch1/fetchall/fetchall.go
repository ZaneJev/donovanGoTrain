// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	f, err := os.Create("fetch.txt")
	if err != nil {
		fmt.Printf("ошибка создания файла: %s", err)
	}

	for range os.Args[1:] {
		_, err = f.WriteString(fmt.Sprintf("%s\n", <-ch)) // receive from channel ch
		if err != nil {
			fmt.Printf("ошибка записи в файл: %s", err)
		}
	}

	_, err = f.WriteString(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))
	if err != nil {
		fmt.Printf("ошибка записи в файл: %s", err)
	}
	defer f.Close()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
