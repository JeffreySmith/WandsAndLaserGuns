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
	var faceCard Card
	var numberCard Card
	g := NewGame()
	if g.P.Health > 5{
		won = true
	}
	for len(g.FDeck.Cards) > 0 && len(g.NDeck.Cards) > 0 {
		faceCard, ok := g.FDeck.Draw()
		if !ok && g.P.Health > 0{
			won = true
			break
		}
		numberCard, ok := g.NDeck.Draw()
		if !ok && g.P.Health > 0 {
			won = true
			break
		}
		

		break
	}
	//Do the game stuff here
	
	result <- won
}
// Returns which cards should be discarded
func CombatRound(player *Player, face, number Card) (bool,bool){
	var SetEffectsOnPlayer bool
	var Roll int
	var suit Suits
	same_suit := SameSuit(face, number)
	total := CalculateDrawTotal(*player, face, number)
	if same_suit {
		
		SetEffectsOnPlayer = true
		suit = face.Suit
	} else {
		suit = number.Suit
	}
	Roll = player.Roll(suit)
	
	if Roll >= total && same_suit {
		player.WinSuitStat(suit)
		return true, true
	} else if Roll >= total {
		player.WinSuitStat(suit)
		return false, true
	} else if same_suit {
		
	} else {
		player.Health -= 1
	}
	
	
	return false,false
}
