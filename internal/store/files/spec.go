package files

type Spec interface {
	Is(*File) bool
}

type IDInSpec struct {
	ID []string
}

func (s IDInSpec) Is(c *File) bool {
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

func (s ProjectIDInSpec) Is(c *File) bool {
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

type FileTypeInSpec struct {
	Type []FileType
}

func (s FileTypeInSpec) Is(f *File) bool {
	for _, val := range s.Type {
		if f.Type == val {
			return true
		}
	}

	return false
}

func FileTypeIn(vals ...FileType) Spec {
	return FileTypeInSpec{Type: vals}
}

type AndSpec struct {
	Specs []Spec
}

func (s AndSpec) Is(c *File) bool {
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

func (s OrSpec) Is(c *File) bool {
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
