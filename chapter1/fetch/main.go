// Fetch выводит ответ на запрос по заданному URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	fetch1()
	fetch2()
	fetch3()
	fetch4()
}

func fetch1() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

// Копирует тело ответа в stdout для экономии памяти
func fetch2() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%d bytes", b)
	}
}

// Проверяет наличие префикса в строке "http://"
func fetch3() {
	for _, url := range os.Args[1:] {
		prefix := "http://"
		if strings.HasPrefix(url, "http://") {
			prefix = ""
		}
		resp, err := http.Get(prefix + url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		//		b, err := io.ReadAll(resp.Body)
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%d bytes", b)
	}
}

// Выводит статус код
func fetch4() {
	for _, url := range os.Args[1:] {
		prefix := "http://"
		if strings.HasPrefix(url, "http://") {
			prefix = ""
		}
		resp, err := http.Get(prefix + url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		resp.Body.Close()
		fmt.Printf("Status: %s", resp.Status)
	}
}
