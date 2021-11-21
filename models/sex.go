package models

type Sex int64

const (
	Male Sex = iota
	Female
	Other
)

func (s Sex) String() string {
	switch s {
	case Male:
		return "male"
	case Female:
		return "female"
	}
	return "other"
}
