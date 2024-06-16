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

func Game(result chan<- bool) {
	var won bool
	game := NewGame()
	if game.P.Health > 5 {
		won = true
	}
	for len(game.FDeck.Cards) > 0 && len(game.NDeck.Cards) > 0 {
		faceCard, ok := game.FDeck.Draw()
		if !ok && game.P.Health > 0 {
			won = true
			break
		}
		numberCard, ok := game.NDeck.Draw()
		if !ok && game.P.Health > 0 {
			won = true
			break
		}
		face := CombatRound(&game.P, faceCard, numberCard)
		game.P.DefeatedPile = append(game.P.DefeatedPile, numberCard)
		if face {
			game.P.DefeatedPile = append(game.P.DefeatedPile, faceCard)
		} else {
			game.FDeck.Cards = append(game.FDeck.Cards, faceCard)
			game.FDeck.Shuffle()
		}

		break
	}

	result <- won
}

// Returns whether the face card should be discarded
func CombatRound(player *Player, face, number Card) bool {

	var Roll int
	var suit Suits
	var card Card
	var skip bool
	same_suit := SameSuit(face, number)
	total := CalculateDrawTotal(*player, face, number)
	if same_suit {
		suit = face.Suit
		card = face
	} else {
		suit = number.Suit
		card = number
	}
	Roll = player.Roll(suit)

	//skip = player.ShouldSkip(face, number, Roll)
	if skip && !same_suit {
		return true
	}
	if Roll >= total && same_suit {
		player.WinSuitStat(card)

		return true
	} else if Roll >= total {
		player.WinSuitStat(card)
		return false
	} else if same_suit {

	} else {
		player.Health -= 1
	}

	return false
}
