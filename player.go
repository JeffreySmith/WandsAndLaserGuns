package wandsandlaserguns

import (
	"math/rand"
	"slices"
)


type Player struct {
	Health int
	Wands int
	Laserguns int
	TurnNumber int
	DiscardPile []Card
	DefeatedPile []Card
	ActiveEffects []Effects
	
}

func RollDie() int {
	return rand.Intn(9)+1
}

func (p *Player) Roll() int {
	var stat int

	if !slices.Contains(p.ActiveEffects, BlockWands) && p.Wands > p.Laserguns {
		stat = p.Wands
	} else if !slices.Contains(p.ActiveEffects, BlockLasers) && p.Laserguns > p.Wands {
		stat = p.Laserguns
	} else if !slices.Contains(p.ActiveEffects, BlockLasers) && !slices.Contains(p.ActiveEffects, BlockWands) && p.Wands == p.Laserguns {
		stat = p.Wands
	}
	
	return (RollDie() + stat)
}

func (p *Player) DiscardStatus()  {
	spades := NumSuits(p.DiscardPile,Spades)
	clubs := NumSuits(p.DiscardPile,Clubs)

	if spades >= 5 {
		if !slices.Contains(p.ActiveEffects, BlockLasers){
			p.ActiveEffects = append(p.ActiveEffects,BlockLasers)
		}
	}
	if clubs >= 5 {
		if !slices.Contains(p.ActiveEffects, BlockWands){
			p.ActiveEffects = append(p.ActiveEffects,BlockWands)
		}
	}
}
