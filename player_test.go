package wandsandlaserguns_test

import (
	"math/rand"
	"testing"

	wl "github.com/JeffreySmith/WandsAndLaserGuns"
	"github.com/google/go-cmp/cmp"
)

func TestDieRoll(t *testing.T) {
	t.Parallel()
	n := wl.RollDie()
	if n < 0 || n > 10 {
		t.Errorf("Expected die roll to be 1d10, got %d", n)
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
		t.Error(cmp.Diff(want, got))
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
		t.Error(cmp.Diff(want, got))
	}

}

func TestPlayerWandRoll(t *testing.T) {

	p := wl.Player{}
	p.Wands = 3
	want := 9

	rand.Seed(1)

	outcome := p.Roll(wl.Diamonds)

	if want != outcome {
		t.Errorf("Want %d, got %d", want, outcome)
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
		t.Errorf("Want %d, got %d", want, outcome)
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

		t.Errorf("Want %d, got %d", want, got)
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
func TestLaserStatLoss(t *testing.T) {
	t.Parallel()
	p := wl.NewPlayer()
	p.Health = 5
	p.Laserguns = 3
	p.AceLoss(wl.Laser)

	want := 1
	got := p.Laserguns
	if p.Health != 5 {
		t.Errorf("Expected health to be 5, got %v", p.Health)
	}
	if want != got {
		t.Errorf("Want %v, got %v", want, got)
	}
	
}
func TestRemove2HealthOnLossToAce(t *testing.T) {
	t.Parallel()
	p := wl.NewPlayer()
	p.Health = 8
	want := 6
	p.AceLoss(wl.Ignore)
	got := p.Health

	if want != got {
		t.Errorf("Want %v, got %v", want, got)
	}
}
func TestWandStatOnLossToAce(t *testing.T) {
	t.Parallel()
	p := wl.NewPlayer()
	p.Health = 4
	p.Wands = 5
	want := 3
	p.AceLoss(wl.Wand)
	got := p.Wands

	if p.Health != 4 {
		t.Errorf("Expected health to be 4, got %v", p.Health)
	}
	if want != got {
		t.Errorf("Want %v, got %v", want, got)
	}
}
func TestOnAceLossWithLowStatsWands(t *testing.T) {
	t.Parallel()
	p := wl.NewPlayer()
	p.Health = 4
	p.Wands = 1
	want := 2
	p.AceLoss(wl.Wand)
	got := p.Health

	if p.Wands != 1 {
		t.Errorf("Expected wand to be 1, got %v", p.Wands)
	}
	if want != got {
		t.Errorf("Want %v, got %v", want, got)
	}
}
func TestOnAceLossWithLowStatsLasers(t *testing.T) {
	t.Parallel()
	p := wl.NewPlayer()
	p.Health = 4
	p.Laserguns = 1
	want := 2
	p.AceLoss(wl.Laser)
	got := p.Health

	if p.Laserguns != 1 {
		t.Errorf("Expected Laserguns to be 1, got %v", p.Laserguns)
	}
	if want != got {
		t.Errorf("Want %v, got %v", want, got)
	}
}
func TestSuccessfulAceEncounterStatIncrease(t *testing.T) {
	t.Parallel()
	p := wl.NewPlayer()
	p.WinAgainstAce()
	wands := p.Wands
	lasers := p.Laserguns

	if wands != 1 || lasers != 1 {
		t.Errorf("Wanted lasers and wands to be 1, got %v and %v", lasers, wands)
	}
}
func TestAceEncounterSuccessRestoreStats(t *testing.T) {
	t.Parallel()
	p := wl.NewPlayer()
	p.Laserguns = 4
	p.ActiveEffects = []wl.Effects{wl.BlockLasers}
	p.WinAgainstAce()
	want := []wl.Effects{}
	got := p.ActiveEffects

	if !cmp.Equal(want,got) {
		t.Error(cmp.Diff(want, got))
	}
	p.ActiveEffects = []wl.Effects{wl.BlockWands}
	p.WinAgainstAce()
	got = p.ActiveEffects

	if !cmp.Equal(want,got) {
		t.Error(cmp.Diff(want, got))
	}

	
}
