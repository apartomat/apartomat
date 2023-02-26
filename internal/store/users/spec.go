package users

type Spec interface {
	Is(*User) bool
}

type IDInSpec struct {
	ID []string
}

func (s IDInSpec) Is(c *User) bool {
	for _, val := range s.ID {
		if c.ID == val {
			return true
		}
	}

	return false
}

func IDIn(vals ...string) Spec {
	return IDInSpec{ID: vals}
}

type EmailInSpec struct {
	Email []string
}

func (s EmailInSpec) Is(c *User) bool {
	for _, val := range s.Email {
		if c.Email == val {
			return true
		}
	}

	return false
}

func EmailIn(vals ...string) Spec {
	return EmailInSpec{Email: vals}
}

type AndSpec struct {
	Specs []Spec
}

func (s AndSpec) Is(c *User) bool {
	for _, spec := range s.Specs {
		if !spec.Is(c) {
			return false
		}
	}

	return true
}

func And(specs ...Spec) Spec {
	return AndSpec{Specs: specs}
}

type OrSpec struct {
	Specs []Spec
}

func (s OrSpec) Is(c *User) bool {
	for _, spec := range s.Specs {
		if spec.Is(c) {
			return true
		}
	}

	return false
}

func Or(specs ...Spec) Spec {
	return OrSpec{Specs: specs}
}
