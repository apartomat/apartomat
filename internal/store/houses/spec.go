package houses

type Spec interface {
	Is(*House) bool
}

// IDInSpec is specification that point House has specified ID
type IDInSpec struct {
	IDs []string
}

func (s IDInSpec) Is(c *House) bool {
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

// ProjectIDInSpec is specification that point House belongs specified Project
type ProjectIDInSpec struct {
	IDs []string
}

func (s ProjectIDInSpec) Is(c *House) bool {
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

type AndSpec struct {
	Specs []Spec
}

func (s AndSpec) Is(c *House) bool {
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

func (s OrSpec) Is(c *House) bool {
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
