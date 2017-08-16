package flags

type ExistentialFlag interface {
	Flag() int64
	IsTrue() bool
	IsFalse() bool
	IsKnown() bool
	Matches(ExistentialFlag) bool
	String() string
}

type PlacetypeFlag interface {
	Matches(PlacetypeFlag) bool
	Placetypes() []string
	String() string
}
