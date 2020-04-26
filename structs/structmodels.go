package structs

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

// constructor like functions for the structs, doesnt have to return a pointer other
func NewSaiyan(name string, power int) *Saiyan {
	return &Saiyan{
		Name:  name,
		Power: power,
	}
}

// type sayan is the receiver of the super function and theyre associated with each other.
func (s *Saiyan) Super() {
	s.Power += 10000
}
