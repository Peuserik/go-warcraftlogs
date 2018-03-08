package types

import "github.com/alexejk/go-warcraftlogs/types/warcraft"

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
  StartTime *int64 `json:"start"`
  EndTime   *int64 `json:"end"`
}

type Fight struct {
	Id                            *int64
	Name                          *string
	Boss                          *int64
	StartTime                     *int64 `json:"start_time"`
	EndTime                       *int64 `json:"end_time"`
	Size                          *int
	Difficulty                    *warcraft.Difficulty
	Kill                          *bool
	Partial                       *int
	BossPercentage                *int
	FightPercentage               *int
	LastPhaseForPercentageDisplay *int
}
