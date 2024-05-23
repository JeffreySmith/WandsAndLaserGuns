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


type Player struct {
	Health int
	Wands int
	Laserguns int
	TurnNumber int
	DiscardPile []Card
	DefeatedPile []Card
	ActiveEffects []Effects
	BlockOnSuit map[Suits] []Effects
}
func NewPlayer() Player{
	p := Player{}
	p.Health = RollDie()
	p.BlockOnSuit = make(map[Suits] []Effects,4)
	return p
}
//Get a slice of effects that are active for that suit.
func (p Player) SuitBlockStatus(suit Suits) []Effects {
	effects, ok := p.BlockOnSuit[suit]
	if !ok {
		return []Effects{}
	}
	return effects
}
//Remove an active effect that only affects one suit.
func (p Player) RemoveSuitEffect(suit Suits, effect Effects) {
	effect_slice, ok := p.BlockOnSuit[suit]
	if !ok {
		return
	}
	for i, e := range effect_slice {
 		if e ==  effect {
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
//Returns a random number between 1-10
func RollDie() int {
	return rand.Intn(9)+1
}
//Returns a random number, factoring in stat levels and blocked stats
func (p *Player) Roll(suit Suits) int {
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
	} else {
		stat = laser
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
