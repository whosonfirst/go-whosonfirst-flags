package flags

type ExistentialFlag interface {
	Flag() int64
	Status() bool
	Confidence() bool
	String() string
}
