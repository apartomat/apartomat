package album_files

import (
	"errors"
	"time"
)

var (
	ErrAlbumFileAlreadyDone = errors.New("album file is done")
)

type Status string

const (
	StatusNew        Status = "NEW"
	StatusInProgress Status = "GENERATING_IN_PROGRESS"
	StatusDone       Status = "GENERATING_DONE"
)

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

func NewAlbumFile(id string, status Status, albumID string, version int) *AlbumFile {
	var (
		now = time.Now()
	)

	return &AlbumFile{
		ID:                  id,
		Status:              status,
		AlbumID:             albumID,
		Version:             version,
		FileID:              nil,
		GeneratingStartedAt: nil,
		GeneratingDoneAt:    nil,
		CreatedAt:           now,
		ModifiedAt:          now,
	}
}

func (f *AlbumFile) Done(ts time.Time) error {
	if f.Status == StatusDone {
		return ErrAlbumFileAlreadyDone
	}

	f.Status = StatusDone
	f.GeneratingDoneAt = &ts

	return nil
}

func (f *AlbumFile) DoneNow() error {
	return f.Done(time.Now())
}

func (f *AlbumFile) Start(ts time.Time) error {
	if f.Status == StatusDone {
		return ErrAlbumFileAlreadyDone
	}

	f.Status = StatusInProgress
	f.GeneratingStartedAt = &ts

	return nil
}

func (f *AlbumFile) StartNow() error {
	return f.Start(time.Now())
}
