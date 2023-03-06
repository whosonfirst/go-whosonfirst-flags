package placetypes

import (
	"testing"
)

func TestNewPlacetypeFlag(t *testing.T) {

	tests := []string{
		"region",
		"region#whosonfirst://",
	}

	for _, n := range tests {

		_, err := NewPlacetypeFlag(n)

		if err != nil {
			t.Fatalf("Failed to create flag for '%s', %v", n, err)
		}
	}
	
}
