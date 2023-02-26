package workspace_users

type Spec interface {
	Is(*WorkspaceUser) bool
}

type IDInSpec struct {
	ID []string
}

func (s IDInSpec) Is(c *WorkspaceUser) bool {
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

type WorkspaceIDInSpec struct {
	WorkspaceID []string
}

func (s WorkspaceIDInSpec) Is(c *WorkspaceUser) bool {
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

type UserIDInSpec struct {
	UserID []string
}

func (s UserIDInSpec) Is(c *WorkspaceUser) bool {
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

type AndSpec struct {
	Specs []Spec
}

func (s AndSpec) Is(c *WorkspaceUser) bool {
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

func (s OrSpec) Is(c *WorkspaceUser) bool {
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
