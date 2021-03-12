package date

import (
	"github.com/whosonfirst/go-whosonfirst-flags"
)

type NullFlag struct {
	flags.DateFlag
}

func NewNullFlag() (flags.DateFlag, error) {
	fl := NullFlag{}
	return &fl, nil
}

func (fl *NullFlag) InnerRange() (*int64, *int64) {
	return nil, nil
}

func (fl *NullFlag) OuterRange() (*int64, *int64) {
	return nil, nil
}

func (fl *EDTFFlag) MatchesAny(others ...flags.DateFlag) bool {
	return false
}

func (fl *EDTFFlag) MatchesAll(others ...flags.DateFlag) bool {
	return false
}

func (fl *NullFlag) String() string {
	return ""
}
