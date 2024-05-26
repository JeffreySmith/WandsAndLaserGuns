package wandsandlaserguns


type GameState struct {
	P Player
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
		P:p,
		FDeck:*f,
		NDeck:*n,
	}
}
