package album_files

import "time"

type AlbumFile struct {
	ID                  string
	Status              Status
	AlbumID             string
	Version             int
	FileID              *string
	GeneratingStartedAt *time.Time
	GeneratingDoneAt    *time.Time
	CreatedAt           time.Time
	ModifiedAt          time.Time
}

type Status string

const (
	StatusNew        Status = "NEW"
	StatusInProgress Status = "GENERATING_IN_PROGRESS"
	StatusDone       Status = "GENERATING_DONE"
)
