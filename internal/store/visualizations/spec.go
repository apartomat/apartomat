package visualizations

type Spec interface {
	Is(*Visualization) bool
}

type IDInSpec struct {
	ID []string
}

func (s IDInSpec) Is(c *Visualization) bool {
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

func (s ProjectIDInSpec) Is(c *Visualization) bool {
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

type RoomIDInSpec struct {
	RoomID []string
}

func (s RoomIDInSpec) Is(c *Visualization) bool {
	for _, val := range s.RoomID {
		if c.ProjectID == val {
			return true
		}
	}

	return false
}

func RoomIDIn(vals ...string) Spec {
	return RoomIDInSpec{RoomID: vals}
}

type StatusInSpec struct {
	Status []VisualizationStatus
}

func (s StatusInSpec) Is(c *Visualization) bool {
	for _, val := range s.Status {
		if c.Status == val {
			return true
		}
	}

	return false
}

func StatusIn(vals ...VisualizationStatus) Spec {
	return StatusInSpec{Status: vals}
}

type StatusNotInSpec struct {
	Status []VisualizationStatus
}

func (s StatusNotInSpec) Is(c *Visualization) bool {
	for _, val := range s.Status {
		if c.Status == val {
			return false
		}
	}

	return true
}

func StatusNotIn(vals ...VisualizationStatus) Spec {
	return StatusNotInSpec{Status: vals}
}

type AndSpec struct {
	Specs []Spec
}

func (s AndSpec) Is(c *Visualization) bool {
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

func (s OrSpec) Is(c *Visualization) bool {
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

func NotDeletedByProjectAndRoom(projectID string, roomID *string) Spec {
	var (
		spec = And(
			ProjectIDIn(projectID),
			StatusNotIn(VisualizationStatusDeleted),
		)
	)

	if roomID != nil {
		spec = And(spec, RoomIDIn(*roomID))
	}

	return spec
}
