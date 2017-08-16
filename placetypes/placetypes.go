package placetypes

import (
	"github.com/go-whosonfirst-flags"
	wof "github.com/go-whosonfirst-placetypes"
	"strings"
)

type PlacetypesFlag struct {
	flags.PlacetypeFlag
	placetypes []*wof.WOFPlacetypes
	names      []string
}

func NewPlacetypesFlag(str_placetypes string) (*PlacetypesFlag, error) {

	require := make([]*wof.WOFPlacetypes, 0)

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
		placetypes: possible,
		names:      names,
	}

	return &f, nil
}

func (f *PlacetypesFlag) Matches(other flag.PlacetypesFlag) bool {

	ours := f.Placetypes()
	theirs := other.Placetypes()

	for _, a := range theirs {

		for _, b := range ours {

			if a == b {
				return true
			}
		}
	}

	return false
}

func (f *PlacetypesFlag) Placetypes() []string {
	return f.names
}

func (f *PlacetypesFlag) String() string {
	return strings.Join(f.Placetypes(), ",")
}
