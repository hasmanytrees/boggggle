package boggle

import (
	"testing"
)

/**
Classic				New
AACIOT				AAEEGN
ABILTY				ABBJOO
ABJMOQu				ACHOPS
ACDEMP				AFFKPS
ACELRS				AOOTTW
ADENVZ				CIMOTU
AHMORS				DEILRX
BIFORX				DELRVY
DENOSW				DISTTY
DKNOTU				EEGHNW
EEFHIY				EEINSU
EGKLUY				EHRTVW
EGINTV				EIOSST
EHINPS				ELRTTY
ELPSTU				HIMNUQu
GILRUW				HLNNRZ
*/

type MockRandomizer struct {
}

func (mr *MockRandomizer) Intn(n int) int {
	return 2
}

func TestRoll(t *testing.T) {
	// setup
	mockRandomizer := MockRandomizer{}
	die := Die{"A", "A", "C", "I", "O", "T"}

	// exercise
	got := die.Roll(&mockRandomizer)

	// assert
	if got != "C" {
		t.Fatalf("Expected rolled value to be in the collection %v but got %s", die, got)
	}
}
