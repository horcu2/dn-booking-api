package models

type ServiceType int64

const (
	Pediatrician ServiceType = iota
	Colonoscopy
	Obgyn
)

func (s ServiceType) String() string {
	switch s {
	case Pediatrician:
		return "pediatrician"
	case Colonoscopy:
		return "colonoscopy"
	}
	return "ob/gyn"
}

type Service struct {
	Id   string `json:"id"`
	Type ServiceType `json:"name"`
}