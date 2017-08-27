package types

type Zone struct {
	Id         *int64
	Name       *string
	Frozen     *bool
	Encounters []*Encounter
	Brackets   []*Bracket
}

type Encounter struct {
	Id   *int64
	Name *string
}

type Bracket struct {
	Id   *int64
	Name *string
}
type Report struct {
	Id        *string
	Title     *string
	Owner     *string
	Zone      *int64
	StartTime *int64
	EndTime   *int64
}
