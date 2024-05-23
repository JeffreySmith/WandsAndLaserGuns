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
	"strconv"
)

type Suits int
type CardValue int

const (
	Hearts Suits = iota
	Diamonds
	Clubs
	Spades
)

type Effects int

const (
	BlockWands Effects = iota
	BlockLasers
	SubtractWands
	SubtractLasers
	AddWands
	AddLasers
	AddHealth
	HeartTally
	DiamondTally
	ClubTally
	SpadeTally
	PlusTwoOrRegain
	Subtract2orHealth
	SubtractHealth
	SkipEncounter
	Sticky
	Ephemeral
	None
)

type Card struct {
	Value  CardValue
	Suit   Suits
	Effect map[string]Effects
	Active bool
}

type Deck struct {
	Cards []Card
}



func (e Effects) String() string {
	switch e {
	case BlockWands:
		return "Block Wands"
	case BlockLasers:
		return "Block Lasers"
	case SubtractWands:
		return "Subtract Wands"
	case SubtractLasers:
		return "Subtract Lasers"
	case AddLasers:
		return "Add Lasers"
	case AddWands:
		return "Add Wands"
	case HeartTally:
		return "Add to Heart Tally"
	case DiamondTally:
		return "Add to Diamond Tally"
	case ClubTally:
		return "Add to Club Tally"
	case SpadeTally:
		return "Add to Spade Tally"
	case PlusTwoOrRegain:
		return "Plus 2 or Regain"
	case AddHealth:
		return "Add Health"
	case Subtract2orHealth:
		return "Subtract 2 or Health"
	case SubtractHealth:
		return "Subtract Health"
	case SkipEncounter:
		return "Skip Encounter"
	case Sticky:
		return "Sticky"
	case Ephemeral:
		return "Ephemeral"
	case None:
		return "None"
	default:
		return "Not implemented"
	}
}

func (c CardValue) String() string {
	out := strconv.Itoa(int(c))
	switch int(c) {
	case 11:
		out = "Jack"
	case 12:
		out = "Queen"
	case 13:
		out = "King"
	case 14:
		out = "Ace"
	}
	return out
}

func (s Suits) String() string {
	switch s {
	case 0:
		return "Hearts"
	case 1:
		return "Diamonds"
	case 2:
		return "Clubs"
	default:
		return "Spades"
	}
}
//Creates a face card and initialises all of their effects.
func initFaceCard(card int, suit Suits) Card {
	var onSuccess Effects
	var onFailure Effects
	duration := Sticky
	new_card := Card{Value: CardValue(card), Suit: suit}
	//Strangely, setting it to make(...,2) seems to be consistently fastest
	new_card.Effect = make(map[string]Effects,2)
	switch suit {
	case Hearts:
		switch card {
		case 11:
			onFailure = HeartTally
			onSuccess = SkipEncounter
		case 12:
			onFailure = BlockWands
			onSuccess = AddWands
		case 13:
			onFailure = BlockLasers
			onSuccess = AddLasers
		}
	case Diamonds:
		switch card {
		case 11:
			onFailure = DiamondTally
			onSuccess = SkipEncounter
		case 12:
			onFailure = BlockWands
			onSuccess = AddWands
		case 13:
			onFailure = BlockLasers
			onSuccess = AddLasers
		}
	case Clubs:
		switch card {
		case 11:
			onFailure = ClubTally
			onSuccess = SkipEncounter
		case 12:
			onFailure = BlockWands
			onSuccess = AddWands
		case 13:
			onFailure = BlockLasers
			onSuccess = AddLasers
		}
	case Spades:
		switch card {
		case 11:
			onFailure = SpadeTally
			onSuccess = SkipEncounter
		case 12:
			onFailure = BlockWands
			onSuccess = AddWands
		case 13:
			onFailure = BlockLasers
			onSuccess = AddLasers
		}
	}
	if card == 14 {
		onFailure = Subtract2orHealth
		onSuccess = PlusTwoOrRegain
		duration = Ephemeral
	}
	new_card.Effect["win"] = onSuccess
	new_card.Effect["lose"] = onFailure
	new_card.Effect["duration"] = duration
	return new_card
}
//Create a new face card only deck
func NewFaceDeck() *Deck {
	deck := Deck{}
	for i := 0; i < 4; i++ {
		for j := 11; j <= 14; j++ {
			deck.Cards = append(deck.Cards, initFaceCard(j, Suits(i)))
		}
	}
	return &deck
}
//Create a new numbers-only deck.
func NewNumberDeck() *Deck {
	deck := Deck{}
	for i := 0; i < 4; i++ {
		for j := 2; j <= 10; j++ {
			deck.Cards = append(deck.Cards, Card{
				Value: CardValue(j),
				Suit:  Suits(i),
			})
		}
	}
	return &deck
}
//Bool here returns false on an empty deck.
//This is the win condition.
func (d *Deck) Draw() (Card, bool) {
	if len(d.Cards) == 0{
		return Card{},false
	}
	c := d.Cards[0]
	d.Cards = d.Cards[1:]

	return c,true
}
//Insert a single card into a deck.
//This only occurs with "boss" cards.
func (d *Deck) InsertCard(c Card) {
	d.Cards = append(d.Cards,c)
	d.Shuffle()

}

func (d *Deck) Shuffle() {
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}
//Count the number of Cards of a suit in a particular card array
//Several rules in the game depend on this
func NumSuits(cards []Card, suit Suits) int {
	var total int
	for _, c := range cards {
		if c.Suit == suit {
			total += 1
		}
	}

	return total
}
//Delete all cards of a particular suit.
//Used after disabling wands/lasers because of the discard pile.
func (d *Deck) RemoveCards(suit Suits){
	for i, c := range d.Cards {
		if c.Suit == suit {
			d.Cards[i] = d.Cards[len(d.Cards)-1]
			d.Cards = d.Cards[:len(d.Cards)-1]
		}
	}
}
//Remove Cards from a card array.
//Primarily used on the Player's discard pile.
func RemoveCards(cards []Card, suit Suits) []Card{
	for i:= len(cards)-1; i>= 0; i-- {
		c := cards[i]
		if c.Suit == suit {
			if i < len(cards){
				cards = append(cards[:i],cards[i+1:]...)

			} else if i == len(cards) -1 {
				cards = cards[:len(cards)-1]
			}
		}
	}
	return cards
}
