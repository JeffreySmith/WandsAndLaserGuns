package main

import (
	"fmt"
	wl "github.com/JeffreySmith/WandsAndLaserGuns"
)

func main(){
	d := wl.NewFaceDeck()
	d.Shuffle()
	e := wl.NewNumberDeck()
	e.Shuffle()
	fmt.Println(d.Cards[0], d.Cards[4])
	fmt.Println(e.Cards[0],e.Cards[4])
}
