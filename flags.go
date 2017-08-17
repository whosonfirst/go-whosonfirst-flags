package flags

type ExistentialFlag interface {
	Flag() int64
	IsTrue() bool
	IsFalse() bool
	IsKnown() bool
	MatchesAny(...ExistentialFlag) bool
	MatchesAll(...ExistentialFlag) bool
	String() string
}

type PlacetypesFlag interface {
	MatchesAny(...PlacetypesFlag) bool
	MatchesAll(...PlacetypesFlag) bool
	Placetypes() []string
	String() string
}
