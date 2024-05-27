package main

import (
	"fmt"
	"sync"

	wl "github.com/JeffreySmith/WandsAndLaserGuns"
)

func main() {
	var wg sync.WaitGroup
	total := 0
	results := make(chan bool,100)
	for i:= 0; i < 100; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			wl.Game(results)
		}()
	}
	go func(){
		wg.Wait()
		close(results)
	}()
	for result := range results {
		if result {
			total += 1
		}
	}
	fmt.Printf("Total wins: %v\n", total)
}
