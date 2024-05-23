package wandsandlaserguns_test

import (
	"math/rand"
	"testing"

	wl "github.com/JeffreySmith/WandsAndLaserGuns"
	"github.com/google/go-cmp/cmp"
)

func TestDieRoll(t *testing.T) {
	t.Parallel()
	n:= wl.RollDie()
	if n < 0 || n > 10 {
		t.Errorf("Expected die roll to be 1d10, got %d",n)
	}
}
func TestCreateNewPlayer(t *testing.T) {
	t.Parallel()
	p := wl.NewPlayer()

	if p.Health < 1 || p.Health > 10 {
		t.Errorf("Expected some health, got %d", p.Health)
	}
}

func TestWandsBlockedOnDiamonds(t *testing.T) {
	t.Parallel()
	p := wl.NewPlayer()

	p.BlockOnSuit[wl.Diamonds] = []wl.Effects{wl.BlockWands}
	want := []wl.Effects{wl.BlockWands}
	got := p.SuitBlockStatus(wl.Diamonds)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want,got))
	}

}
func TestWandsUnBlockedOnDiamonds(t *testing.T) {
	t.Parallel()
	p := wl.NewPlayer()

	p.BlockOnSuit[wl.Diamonds] = []wl.Effects{wl.BlockWands}
	want := []wl.Effects{}
	p.RemoveSuitEffect(wl.Diamonds, wl.BlockWands)
	got := p.SuitBlockStatus(wl.Diamonds)

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want,got))
	}

}

func TestPlayerWandRoll(t *testing.T) {

	p := wl.Player{}
	p.Wands = 3
	want := 9

	rand.Seed(1)

	outcome := p.Roll(wl.Diamonds)

	if want != outcome {
		t.Errorf("Want %d, got %d",want,outcome)
	}
}
func TestPlayerLaserRoll(t *testing.T) {
	
	p := wl.Player{}
	p.Laserguns = 3
	want := 9
	rand.Seed(1)


	outcome := p.Roll(wl.Diamonds)

	if want != outcome {
		t.Log("This test fails *sometimes*")
		t.Errorf("Want %d, got %d",want,outcome)
	}
}
func TestLasersAndWandsDisabled(t *testing.T) {

	rand.Seed(1)
	p := wl.Player{}
	p.Laserguns = 5
	p.Wands = 5
	p.ActiveEffects = []wl.Effects{wl.BlockLasers, wl.BlockWands}
	want := 6
	got := p.Roll(wl.Diamonds)
	if got != want {
		
		t.Errorf("Want %d, got %d",want,got)
	}
}
func TestPlayerBlockedOnSuit(t *testing.T) {
	rand.Seed(1)
	p := wl.NewPlayer()
	p.Laserguns = 12
	p.Wands = 0
	p.BlockOnSuit[wl.Diamonds] = []wl.Effects{wl.BlockLasers}
	want := 7
	got := p.Roll(wl.Diamonds)

	if want != got {
		t.Log(got)
		t.Errorf("Want %v, got %v", want, got)
	}
}
