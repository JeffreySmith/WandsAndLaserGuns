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

	if cmp.Equal(deck,unshuffled){
		t.Error(cmp.Diff(unshuffled, deck))
	}
}

func TestShufflingNumberCardDeck(t *testing.T) {
	t.Parallel()
	rand.New(rand.NewSource(1))
	deck := wl.NewNumberDeck()
	deck.Shuffle()
	unshuffled := wl.NewNumberDeck()

	if cmp.Equal(unshuffled, deck) {
		t.Error(cmp.Diff(deck,unshuffled))
	}
}
