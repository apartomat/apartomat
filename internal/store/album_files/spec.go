package album_files

type Spec interface {
	Is(*AlbumFile) bool
}

//

type IDInSpec struct {
	ID []string
}

func (s IDInSpec) Is(c *AlbumFile) bool {
	for _, val := range s.ID {
		if c.ID == val {
			return true
		}
	}

	return false
}

func IDIn(vals ...string) IDInSpec {
	return IDInSpec{ID: vals}
}

//

type AlbumIDInSpec struct {
	AlbumID []string
}

func (s AlbumIDInSpec) Is(c *AlbumFile) bool {
	for _, val := range s.AlbumID {
		if c.AlbumID == val {
			return true
		}
	}

	return false
}

func AlbumIDIn(vals ...string) Spec {
	return AlbumIDInSpec{AlbumID: vals}
}

//

type VersionInSpec struct {
	Version []int
}

func (s VersionInSpec) Is(c *AlbumFile) bool {
	for _, val := range s.Version {
		if c.Version == val {
			return true
		}
	}

	return false
}

func VersionIn(vals ...int) Spec {
	return VersionInSpec{Version: vals}
}

//

type VersionGteSpec struct {
	Version int
}

func (s VersionGteSpec) Is(f *AlbumFile) bool {
	return f.Version >= s.Version
}

func VersionGte(val int) Spec {
	return VersionGteSpec{Version: val}
}

//

type AndSpec struct {
	Specs []Spec
}

func (s AndSpec) Is(c *AlbumFile) bool {
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

func (s OrSpec) Is(c *AlbumFile) bool {
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
