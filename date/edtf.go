package date

import (
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/parser"
	"github.com/whosonfirst/go-whosonfirst-flags"
)

const CONTAINS int = 0
const INTERSECTS int = 1

const INNER int = 0
const OUTER int = 1

type EDTFFlag struct {
	flags.DateFlag
	date     *edtf.EDTFDate
	mode     int
	boundary int
}

func NewEDTFFlagsArray(names ...string) ([]flags.DateFlag, error) {

	pt_flags := make([]flags.DateFlag, 0)

	for _, name := range names {

		fl, err := NewEDTFFlag(name)

		if err != nil {
			return nil, err
		}

		pt_flags = append(pt_flags, fl)
	}

	return pt_flags, nil
}

func NewEDTFFlag(edtf_str string) (flags.DateFlag, error) {

	d, err := parser.ParseString(edtf_str)

	if err != nil {
		return nil, err
	}

	fl := EDTFFlag{
		date:     d,
		mode:     CONTAINS,
		boundary: OUTER,
	}

	return &fl, nil
}

func (fl *EDTFFlag) InnerRange() (*int64, *int64) {

	start_ts := fl.date.Start.Upper.Timestamp
	end_ts := fl.date.End.Lower.Timestamp

	return fl.describeRange(start_ts, end_ts)
}

func (fl *EDTFFlag) OuterRange() (*int64, *int64) {

	start_ts := fl.date.Start.Lower.Timestamp
	end_ts := fl.date.End.Upper.Timestamp

	return fl.describeRange(start_ts, end_ts)
}

func (fl *EDTFFlag) describeRange(start_ts *edtf.Timestamp, end_ts *edtf.Timestamp) (*int64, *int64) {

	if start_ts != nil && end_ts != nil {
		start := start_ts.Unix()
		end := end_ts.Unix()
		return &start, &end
	}

	if start_ts != nil {
		start := start_ts.Unix()
		return &start, nil
	}

	if end_ts != nil {
		end := end_ts.Unix()
		return nil, &end
	}

	return nil, nil
}

func (fl *EDTFFlag) MatchesAny(others ...flags.DateFlag) bool {

	for _, o := range others {

		if fl.matches(o) {
			return true
		}
	}

	return false
}

func (fl *EDTFFlag) MatchesAll(others ...flags.DateFlag) bool {

	matches := 0

	for _, o := range others {

		if fl.matches(o) {
			matches += 1
		}

	}

	if matches == len(others) {
		return true
	}

	return false
}

func (fl *EDTFFlag) matches(o flags.DateFlag) bool {

	switch fl.boundary {
	case INNER:
		switch fl.mode {
		case INTERSECTS:
			return fl.intersectsInner(o)
		default:
			return fl.intersectsOuter(o)
		}
	default:
		switch fl.mode {
		case INTERSECTS:
			return fl.containsInner(o)
		default:
			return fl.containsOuter(o)
		}
	}
}

func (fl *EDTFFlag) containsInner(o flags.DateFlag) bool {

	start, end := fl.InnerRange()
	o_start, o_end := o.InnerRange()

	if start == nil || end == nil {
		return false
	}

	if o_start == nil || o_end == nil {
		return false
	}

	return fl.contains(*start, *end, *o_start, *o_end)
}

func (fl *EDTFFlag) containsOuter(o flags.DateFlag) bool {

	start, end := fl.OuterRange()
	o_start, o_end := o.OuterRange()

	if start == nil || end == nil {
		return false
	}

	if o_start == nil || o_end == nil {
		return false
	}

	return fl.contains(*start, *end, *o_start, *o_end)
}

func (fl *EDTFFlag) contains(start int64, end int64, o_start int64, o_end int64) bool {

	if start > o_start {
		return false
	}

	if end < o_end {
		return false
	}

	return true
}

func (fl *EDTFFlag) intersectsInner(o flags.DateFlag) bool {

	start, end := fl.InnerRange()
	o_start, o_end := o.InnerRange()

	if start == nil || end == nil {
		return false
	}

	if o_start == nil || o_end == nil {
		return false
	}

	return fl.intersects(*start, *end, *o_start, *o_end)
}

func (fl *EDTFFlag) intersectsOuter(o flags.DateFlag) bool {

	start, end := fl.OuterRange()
	o_start, o_end := o.OuterRange()

	if start == nil || end == nil {
		return false
	}

	if o_start == nil || o_end == nil {
		return false
	}

	return fl.intersects(*start, *end, *o_start, *o_end)
}

func (fl *EDTFFlag) intersects(start int64, end int64, o_start int64, o_end int64) bool {

	if o_start > end {
		return false
	}

	if o_end < start {
		return false
	}

	return true
}

func (fl *EDTFFlag) String() string {
	return fl.date.EDTF
}
