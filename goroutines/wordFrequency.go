package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

// https://github.com/cutajarj/ConcurrentProgrammingWithGo

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	frequency := make(map[string]uint8)

	wg.Add(200)
	for i := 1_000; i < 1_200; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countWordsFrom(url, frequency, &wg, &mutex)
	}
	wg.Wait()

	for k, v := range frequency {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func countWordsFrom(url string, frequency map[string]uint8, wg *sync.WaitGroup, mutex *sync.Mutex) {
	resp, _ := http.Get(url)
	defer func() {
		_ = resp.Body.Close()
		wg.Done()
	}()

	mutex.Lock()
	wordRegex := regexp.MustCompile(`[a-zA-Z]+`)
	body, _ := io.ReadAll(resp.Body)
	for _, word := range wordRegex.FindAllString(string(body), -1) {
		wordLower := strings.ToLower(word)
		frequency[wordLower] += 1
	}
	mutex.Unlock()

	fmt.Printf("Completed: %s\n", url)
}
