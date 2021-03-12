package date

import (
	_ "fmt"
	"testing"
)

func TestEDTFDate(t *testing.T) {

	d1_str := "2020-03-10/2020-04-29"
	d2_str := "2020-04-01"
	d3_str := "2020-05-01"

	d1, err := NewEDTFFlag(d1_str)

	if err != nil {
		t.Fatalf("Failed to create date flag for '%s', %v", d1_str, err)
	}

	if d1.String() != d1_str {
		t.Fatalf("Invalid string value. Expected '%s' but got '%s'", d1_str, d1.String())
	}

	d2, err := NewEDTFFlag(d2_str)

	if err != nil {
		t.Fatalf("Failed to create date flag for '%s', %v", d2_str, err)
	}

	if !d1.MatchesAll(d2) {
		t.Fatalf("Match fail: '%s' should match '%s'", d2_str, d1_str)
	}

	d3, err := NewEDTFFlag(d3_str)

	if err != nil {
		t.Fatalf("Failed to create date flag for '%s', %v", d3_str, err)
	}

	if d1.MatchesAll(d3) {
		t.Fatalf("Match fail: '%s' should not match '%s'", d3_str, d1_str)
	}

	if !d1.MatchesAny(d2, d3) {
		t.Fatalf("Match fail: '%s' should match at least one of '%s', '%s'", d1_str, d2_str, d3_str)
	}

	nd, err := NewNullDateFlag()

	if err != nil {
		t.Fatalf("Failed to create null date flag, %v", err)
	}

	if d1.MatchesAny(nd) {
		t.Fatalf("Match fail: '%s' should not match null date", d1_str)
	}

	if nd.MatchesAny(d2) {
		t.Fatalf("Match fail: null date should not match '%s'", d2_str)
	}
}
