package placetypes

import (
	"github.com/whosonfirst/go-whosonfirst-flags"
)

type NullFlag struct {
	flags.PlacetypesFlag
}

func NewNullFlag() (*NullFlag, error) {

	f := NullFlag{}
	return &f, nil
}

func (f *NullFlag) MatchesAny(others ...flags.PlacetypesFlag) bool {
	return true
}

func (f *NullFlag) MatchesAll(others ...flags.PlacetypesFlag) bool {
	return true
}

func (f *NullFlag) Placetypes() []string {
	return []string{}
}

func (f *NullFlag) String() string {
	return "NULL"
}
