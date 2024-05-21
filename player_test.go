package wandsandlaserguns_test

import (
	"math/rand"
	"testing"

	wl "github.com/JeffreySmith/WandsAndLaserGuns"
	//"github.com/google/go-cmp/cmp"
)

func TestDieRoll(t *testing.T) {
	t.Parallel()
	n:= wl.RollDie()
	if n < 0 || n > 10 {
		t.Errorf("Expected die roll to be 1d10, got %d",n)
	}
}
//FIXME:This test fails *sometimes*
func TestPlayerWandRoll(t *testing.T) {
	t.Parallel()
	p := wl.Player{}
	p.Wands = 3
	want := 9

	rand.Seed(1)

	outcome := p.Roll()

	if want != outcome {
		t.Log("This test fails sometimes")
		t.Errorf("Want %d, got %d",want,outcome)
	}
}
func TestPlayerLaserRoll(t *testing.T) {
	t.Parallel()
	p := wl.Player{}
	p.Laserguns = 3
	want := 9
	rand.Seed(1)


	outcome := p.Roll()

	if want != outcome {
		t.Log("This test fails *sometimes*")
		t.Errorf("Want %d, got %d",want,outcome)
	}
}
func TestLasersAndWandsDisabled(t *testing.T) {
	t.Parallel()

	rand.Seed(1)
	p := wl.Player{}
	p.Laserguns = 5
	p.Wands = 5
	p.ActiveEffects = []wl.Effects{wl.BlockLasers, wl.BlockWands}
	want := 6
	got := p.Roll()
	if got != want {
		t.Log("This test might fail if something is up with the rand seed")
		t.Errorf("Want %d, got %d",want,got)
	}
}
