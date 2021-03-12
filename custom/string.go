package custom

import (
	"github.com/whosonfirst/go-whosonfirst-flags"
)

type CustomStringFlag struct {
	flags.CustomFlag
	value string
}

func NewCustomCustomStringFlag(value string) (flags.CustomFlag, error) {

	fl := &CustomStringFlag{
		value: value,
	}

	return fl, nil
}

func (fl *CustomStringFlag) MatchesAny(others ...flags.CustomFlag) bool {

	for _, o := range others {

		if fl.String() == o.String() {
			return true
		}

	}

	return false
}

func (fl *CustomStringFlag) MatchesAll(others ...flags.CustomFlag) bool {

	matches := 0

	for _, o := range others {

		if fl.String() == o.String() {
			matches += 1
		}

	}

	if matches == len(others) {
		return true
	}

	return false
}

func (fl *CustomStringFlag) String() string {
	return fl.value
}
