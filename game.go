package wandsandlaserguns

type GameState struct {
	P     Player
	FDeck Deck
	NDeck Deck
}

func NewGame() GameState {
	p := NewPlayer()
	f := NewFaceDeck()
	n := NewNumberDeck()
	f.Shuffle()
	n.Shuffle()
	return GameState{
		P:     p,
		FDeck: *f,
		NDeck: *n,
	}
}

func Game(result chan <- bool) {
	var won bool
	g := NewGame()
	if g.P.Health > 5{
		won = true
	}
	//Do the game stuff here
	
	result <- won
}
