package date

import (
	"testing"
)

func TestNullDate(t *testing.T) {

	n1, err := NewNullDateFlag()

	if err != nil {
		t.Fatalf("Failed to create null date flag, %v", err)
	}

	n2, err := NewNullDateFlag()

	if err != nil {
		t.Fatalf("Failed to create null date flag, %v", err)
	}

	if !n1.MatchesAny(n2) {
		t.Fatalf("Match fail: null date should match null date")
	}
}
