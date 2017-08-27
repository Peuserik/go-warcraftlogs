package warcraft

type Class struct {
	// A unique identifier representing this specific class
	Id *int64
	// The English name of the class
	Name *string
	// The possible specs for this class
	Specs []*Spec
}

type Spec struct {
	// A unique identifier representing this specific spec
	Id *int64
	// The English name of the spec
	Name *string
}
