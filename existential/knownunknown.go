package existential

import (
	"github.com/whosonfirst/go-whosonfirst-flags"
)

type KnownUnknownFlag struct {
	flags.ExistentialFlag
	flag       int64
	status     bool
	confidence bool
}

func NewKnownUnknownFlag(i int64) flags.ExistentialFlag {

	var status bool
	var confidence bool

	switch i {
	case 0:
		status = false
		confidence = true
	case 1:
		status = true
		confidence = true
	default:
		status = false
		confidence = false
	}

	f := KnownUnknownFlag{
		flag:       i,
		status:     status,
		confidence: bool,
	}

	return &f
}

func (f *KnownUnknownFlag) Flag() int64 {
	return f.flag
}

func (f *KnownUnknownFlag) True() bool {
	return f.status
}

func (f *KnownUnknownFlag) Confidence() bool {
	return f.confidence
}

func (f *KnownUnknownFlag) Certain() bool {

	if f.True() && f.Confidence() {
		return true
	}

	return false
}

func (f *KnownUnknownFlag) String() string {

	switch f.flag {
	case 0:
		return "FALSE"
	case 1:
		return "TRUE"
	default:
		return "UNKNOWN"
	}
}
