package wandsandlaserguns_test

import (
	"math/rand"
	"testing"

	wl "github.com/JeffreySmith/WandsAndLaserGuns"
	"github.com/google/go-cmp/cmp"
)

func TestShufflingFaceCardDeck(t *testing.T) {
	t.Parallel()
	rand.New(rand.NewSource(1))
	deck := wl.NewFaceDeck()
	deck.Shuffle()
	unshuffled := wl.NewFaceDeck()

	if cmp.Equal(deck, unshuffled) {
		t.Error(cmp.Diff(unshuffled, deck))
	}
}

func TestShufflingNumberCardDeck(t *testing.T) {
	t.Parallel()
	rand.New(rand.NewSource(1))
	deck := wl.NewNumberDeck()
	deck.Shuffle()
	unshuffled := wl.NewNumberDeck()
	//This shouldn't ever happen
	//Odds are that in human history, no one has achieved this yet
	if cmp.Equal(unshuffled, deck) {
		t.Error(cmp.Diff(deck, unshuffled))
	}
}
func TestSuitCountInDeck(t *testing.T) {
	t.Parallel()
	d := wl.NewFaceDeck()
	want := 4
	got := wl.NumSuits(d.Cards, wl.Hearts)

	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}
func TestAddCardToDeck(t *testing.T) {
	t.Parallel()
	d := wl.NewFaceDeck()
	//Total number of cards in Face Card deck
	want := 16
	c, ok := d.Draw()
	if !ok {
		t.Errorf("Expected cards, got %v", ok)
	}
	d.InsertCard(c)
	got := len(d.Cards)
	if got != want {
		t.Errorf("Got %v cards, wanted %v", got, want)
	}
}
func TestDrawFromDeck(t *testing.T) {
	t.Parallel()

	d := wl.NewFaceDeck()
	//16 total face cards - 1
	want := 15
	_, ok := d.Draw()
	if !ok {
		t.Errorf("Expected a card, got none")
	}
	got := len(d.Cards)

	if got != want {
		t.Errorf("Want %d, got %d", want, got)
	}
}

func TestRemoveAllofClubsFromDeck(t *testing.T) {
	t.Parallel()

	d := wl.NewFaceDeck()
	want := 12
	d.RemoveCards(wl.Clubs)
	got := len(d.Cards)

	if want != got {
		t.Errorf("Want %v, got %v", want, got)
	}
}

func TestPlayerDiscardPileSetsEffects(t *testing.T) {
	t.Parallel()

	d := wl.NewNumberDeck()
	p := wl.Player{DiscardPile: d.Cards}
	want := []wl.Effects{wl.BlockLasers, wl.BlockWands}
	p.DiscardStatus()

	got := p.ActiveEffects

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestRemoveCardsFromSlice(t *testing.T) {
	t.Parallel()
	d := wl.NewFaceDeck()

	want := 8
	d.Cards = wl.RemoveCards(d.Cards, wl.Spades)
	d.Cards = wl.RemoveCards(d.Cards, wl.Clubs)
	got := len(d.Cards)

	if want != got {
		t.Errorf("Want %v, got %v", want, got)
	}
}

func Benchmark_CreateFaceDeck(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = wl.NewFaceDeck()
	}
}
