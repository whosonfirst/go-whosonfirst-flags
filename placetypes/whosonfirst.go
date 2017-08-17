package placetypes

import (
	"github.com/whosonfirst/go-whosonfirst-flags"
	wof "github.com/whosonfirst/go-whosonfirst-placetypes"
	"strings"
)

type PlacetypesFlag struct {
	flags.PlacetypesFlag
	required []*wof.WOFPlacetype
	names    []string
}

func NewPlacetypesFlag(str_placetypes string) (*PlacetypesFlag, error) {

	require := make([]*wof.WOFPlacetype, 0)
	names := make([]string, 0)

	for _, p := range strings.Split(str_placetypes, ",") {

		p = strings.Trim(p, " ")

		pt, err := wof.GetPlacetypeByName(p)

		if err != nil {
			return nil, err
		}

		require = append(require, pt)
		names = append(names, pt.Name)
	}

	f := PlacetypesFlag{
		required: require,
		names:    names,
	}

	return &f, nil
}

func (f *PlacetypesFlag) MatchesAny(others ...flags.PlacetypesFlag) bool {

	ours := f.Placetypes()

	for _, o := range others {

		theirs := o.Placetypes()

		for _, a := range theirs {

			for _, b := range ours {

				if a == b {
					return true
				}
			}
		}
	}

	return false
}

func (f *PlacetypesFlag) MatchesAll(others ...flags.PlacetypesFlag) bool {

	ours := f.Placetypes()
	matches := 0

	for _, o := range others {

		theirs := o.Placetypes()

		for _, a := range theirs {

			for _, b := range ours {

				if a == b {
					matches += 1
				}
			}
		}
	}

	if matches == len(others) {
		return true
	}

	return false
}

func (f *PlacetypesFlag) Placetypes() []string {
	return f.names
}

func (f *PlacetypesFlag) String() string {
	return strings.Join(f.Placetypes(), ",")
}
