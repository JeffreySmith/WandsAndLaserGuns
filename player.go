/*
   BSD 3-Clause License

Copyright (c) 2024, Jeffrey Smith

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its
   contributors may be used to endorse or promote products derived from
   this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package wandsandlaserguns

import (
	"math/rand"
	"slices"
)

type Stat int

const (
	Wand Stat = iota
	Laser
	Health
	Ignore
)

type Player struct {
	Health        int
	MaxHealth     int
	Wands         int
	Laserguns     int
	TurnNumber    int
	ActiveStat    Stat
	DiscardPile   []Card
	DefeatedPile  []Card
	ActiveEffects []Effects
	BlockOnSuit   map[Suits][]Effects
	SkipTokens    int
}

func (s Stat) String() string {
	switch s {
	case Wand:
		return "Wand"
	case Laser:
		return "Laser"
	case Ignore:
		return "No stat"
	default:
		return "Shouldn't happen"
	}
}
func NewPlayer() Player {
	p := Player{}
	p.MaxHealth = RollDie()
	p.Health = p.MaxHealth
	p.BlockOnSuit = make(map[Suits][]Effects, 4)
	return p
}

func (p *Player) RemoveActiveEffect(effect Effects) {
	for i := len(p.ActiveEffects) - 1; i >= 0; i-- {

		if p.ActiveEffects[i] == effect {
			if i == len(p.ActiveEffects)-1 {
				p.ActiveEffects = p.ActiveEffects[:len(p.ActiveEffects)-1]
			} else if i < len(p.ActiveEffects) {
				p.ActiveEffects = append(p.ActiveEffects[:i], p.ActiveEffects[i+1:]...)

			}
		}
	}

}

// Get a slice of effects that are active for that suit.
func (p Player) SuitBlockStatus(suit Suits) []Effects {
	effects, ok := p.BlockOnSuit[suit]
	if !ok {
		return []Effects{}
	}
	return effects
}

// Remove an active effect that only affects one suit.
func (p Player) RemoveSuitEffect(suit Suits, effect Effects) {
	effect_slice, ok := p.BlockOnSuit[suit]
	if !ok {
		return
	}
	for i, e := range effect_slice {
		if e == effect {
			if len(effect_slice) == 1 {
				effect_slice = []Effects{}
			} else if i == 0 {
				effect_slice = []Effects{effect_slice[1]}
			} else if i == 1 {
				effect_slice = []Effects{effect_slice[0]}
			}
		}
	}
	p.BlockOnSuit[suit] = effect_slice
}

// Returns a random number between 1-10
func RollDie() int {
	return rand.Intn(9) + 1
}

// Returns a random number, factoring in stat levels and blocked stats
func (p *Player) Roll(suit Suits) int {
	p.ActiveStat = Ignore
	var stat int
	var wand, laser int
	effects := p.SuitBlockStatus(suit)
	wand = p.Wands
	laser = p.Laserguns

	if slices.Contains(p.ActiveEffects, BlockWands) || slices.Contains(effects, BlockWands) {
		wand = 0
	}
	if slices.Contains(p.ActiveEffects, BlockLasers) || slices.Contains(effects, BlockLasers) {
		laser = 0
	}

	if wand >= laser {
		stat = wand
		p.ActiveStat = Wand
	} else {
		stat = laser
		p.ActiveStat = Laser
	}
	if stat == 0 {
		p.ActiveStat = Ignore
	}
	return (RollDie() + stat)
}

func (p *Player) DiscardStatus() {
	spades := NumSuits(p.DiscardPile, Spades)
	clubs := NumSuits(p.DiscardPile, Clubs)

	if spades >= 5 {
		if !slices.Contains(p.ActiveEffects, BlockLasers) {
			p.ActiveEffects = append(p.ActiveEffects, BlockLasers)
			p.DiscardPile = RemoveCards(p.DiscardPile, Spades)
		}
	}
	if clubs >= 5 {
		if !slices.Contains(p.ActiveEffects, BlockWands) {
			p.ActiveEffects = append(p.ActiveEffects, BlockWands)
			p.DiscardPile = RemoveCards(p.DiscardPile, Clubs)
		}
	}
}

func (p *Player) AceLoss(stat Stat) {
	if stat == Wand {
		if p.Wands >= 2 {
			p.Wands -= 2
		} else {
			p.Health -= 2
		}
	} else if stat == Laser {
		if p.Laserguns >= 2 {
			p.Laserguns -= 2
		} else {
			p.Health -= 2
		}
	} else {
		p.Health -= 2
	}
	return
}
func (p *Player) WinAgainstAce() {
	if slices.Contains(p.ActiveEffects, BlockLasers) && p.Laserguns >= p.Wands {
		p.RemoveActiveEffect(BlockLasers)

	} else if slices.Contains(p.ActiveEffects, BlockWands) {
		p.RemoveActiveEffect(BlockWands)
	} else {
		p.AddStat(Laser)
		p.AddStat(Wand)
	}
}

// Add one to the player's stat.
func (p *Player) AddStat(stat Stat) {
	if stat == Health {
		if p.MaxHealth == p.Health {
			p.MaxHealth += 1
		} else {
			p.Health += 1
		}
	} else if stat == Laser {
		p.Laserguns += 1
	} else if stat == Wand {
		p.Wands += 1
	}
}

func (p *Player) WinSuitStat(card Card) {
	switch card.Suit {
	case Diamonds:
		p.DefeatedPile = append(p.DefeatedPile, card)
	case Hearts:
		if p.Health < p.MaxHealth {
			p.Health += 1
		} else {
			p.MaxHealth += 1
		}
	case Clubs:
		p.AddStat(Wand)
	case Spades:
		p.AddStat(Laser)
	}
}

func (p Player) TallyEffect(suit Suits) bool {
	var active bool

	switch suit {
	case Diamonds:
		if slices.Contains(p.ActiveEffects, DiamondTally) {
			active = true
		}
	case Clubs:
		if slices.Contains(p.ActiveEffects, ClubTally) {
			active = true
		}
	case Hearts:
		if slices.Contains(p.ActiveEffects, HeartTally) {
			active = true
		}
	case Spades:
		if slices.Contains(p.ActiveEffects, SpadeTally) {
			active = true
		}
	}
	return active
}

func (p Player) NumberOfDefeated(suit Suits) int {
	var count int
	for _, c := range p.DefeatedPile {
		if c.Suit == suit {
			count++
		}
	}
	return count
}

func (p *Player) CheckDiscardPileForToken() {
	count := p.NumberOfDefeated(Diamonds)
	if count >= 10 {
		p.DefeatedPile = RemoveCardsFinite(p.DefeatedPile, 10, Diamonds)
		p.SkipTokens += 1
	}
}

func (p *Player) AddEffect(effect Effects) {
	if !slices.Contains(p.ActiveEffects, effect){
		p.ActiveEffects = append(p.ActiveEffects, effect)
	}
}
