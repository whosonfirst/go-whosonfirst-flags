package existential

import (
	"github.com/whosonfirst/go-whosonfirst-flags"
)

type NullFlag struct {
	flags.ExistentialFlag
}

func NewNullFlag() flags.ExistentialFlag {

	n := NullFlag{}
	return &n
}

func (f *KnownUnknownFlag) Flag() int64 {
	return -999
}

func (f *KnownUnknownFlag) IsTrue() bool {
	return false
}

func (f *KnownUnknownFlag) IsFalse() bool {
	return false
}

func (f *KnownUnknownFlag) IsKnown() bool {
	return false
}

func (f *KnownUnknownFlag) String() string {
	return "NULL"
}
