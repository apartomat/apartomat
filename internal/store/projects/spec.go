package projects

type Spec interface {
	Is(*Project) bool
}

//
type IDInSpec struct {
	ID []string
}

func (s IDInSpec) Is(c *Project) bool {
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
type WorkspaceIDInSpec struct {
	WorkspaceID []string
}

func (s WorkspaceIDInSpec) Is(c *Project) bool {
	for _, val := range s.WorkspaceID {
		if c.WorkspaceID == val {
			return true
		}
	}

	return false
}

func WorkspaceIDIn(vals ...string) Spec {
	return WorkspaceIDInSpec{WorkspaceID: vals}
}

//
type StatusInSpec struct {
	Status []Status
}

func (s StatusInSpec) Is(c *Project) bool {
	for _, val := range s.Status {
		if c.Status == val {
			return true
		}
	}

	return false
}

func StatusIn(vals ...Status) Spec {
	return StatusInSpec{Status: vals}
}

//
type AndSpec struct {
	Specs []Spec
}

func (s AndSpec) Is(c *Project) bool {
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

func (s OrSpec) Is(c *Project) bool {
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
