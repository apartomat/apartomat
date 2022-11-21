package workspaces

type Spec interface {
	Is(*Workspace) bool
}

//
type IDInSpec struct {
	ID []string
}

func (s IDInSpec) Is(c *Workspace) bool {
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

//
type UserIDInSpec struct {
	UserID []string
}

func (s UserIDInSpec) Is(c *Workspace) bool {
	for _, val := range s.UserID {
		if c.UserID == val {
			return true
		}
	}

	return false
}

func UserIDIn(vals ...string) Spec {
	return UserIDInSpec{UserID: vals}
}

//
type AndSpec struct {
	Specs []Spec
}

func (s AndSpec) Is(c *Workspace) bool {
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

//
type OrSpec struct {
	Specs []Spec
}

func (s OrSpec) Is(c *Workspace) bool {
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
