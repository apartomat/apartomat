package projectpage

type Spec interface {
	Is(*ProjectPage) bool
}

type IDInSpec struct {
	ID []string
}

func (s IDInSpec) Is(c *ProjectPage) bool {
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

type ProjectIDInSpec struct {
	ProjectID []string
}

func (s ProjectIDInSpec) Is(c *ProjectPage) bool {
	for _, val := range s.ProjectID {
		if c.ProjectID == val {
			return true
		}
	}

	return false
}

func ProjectIDIn(vals ...string) Spec {
	return ProjectIDInSpec{ProjectID: vals}
}

type StatusIsSpec struct {
	Status Status
}

func (s StatusIsSpec) Is(c *ProjectPage) bool {
	return c.Status == s.Status
}

func Public() Spec {
	return StatusIsSpec{Status: StatusPublic}
}

func NotPublic() Spec {
	return StatusIsSpec{Status: StatusNotPublic}
}

type AndSpec struct {
	Specs []Spec
}

func (s AndSpec) Is(c *ProjectPage) bool {
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

func (s OrSpec) Is(c *ProjectPage) bool {
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
