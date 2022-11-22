package contacts

// Spec is a specification for Contact
type Spec interface {
	Is(*Contact) bool
}

// AndSpec is a conjunction of specifications
type AndSpec struct {
	Specs []Spec
}

func (s AndSpec) Is(c *Contact) bool {
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

// OrSpec is a disjunction of specifications
type OrSpec struct {
	Specs []Spec
}

func (s OrSpec) Is(c *Contact) bool {
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

// IDInSpec is specification that point Contact has specified IDIn
type IDInSpec struct {
	IDs []string
}

func (s IDInSpec) Is(c *Contact) bool {
	for _, id := range s.IDs {
		if c.ID == id {
			return true
		}
	}

	return false
}

func IDIn(ids ...string) Spec {
	return IDInSpec{IDs: ids}
}

// ProjectIDInSpec is specification that point Contact belongs specified Project
type ProjectIDInSpec struct {
	IDs []string
}

func (s ProjectIDInSpec) Is(c *Contact) bool {
	for _, id := range s.IDs {
		if c.ProjectID == id {
			return true
		}
	}

	return false
}

func ProjectIDIn(ids ...string) Spec {
	return ProjectIDInSpec{IDs: ids}
}
