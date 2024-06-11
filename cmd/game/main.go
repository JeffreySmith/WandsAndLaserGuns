package main

import (
	"fmt"
	"sync"

	wl "github.com/JeffreySmith/WandsAndLaserGuns"
)

func main() {
	var wg sync.WaitGroup
	var percent float32
	var total int
	number_of_games := 125
	results := make(chan bool,number_of_games)
	for i:= 0; i < number_of_games; i++ {
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
	percent = float32(total)/float32(number_of_games) * 100.0
	
	fmt.Printf("Total wins: %v out of %v\n", total, number_of_games)
	fmt.Printf("Win percentage: %0.1f%%\n", percent)
}
