package files

import (
	"time"
)

type File struct {
	ID         string
	Name       string
	URL        string
	Type       FileType
	MimeType   string
	Size       int64
	CreatedAt  time.Time
	ModifiedAt time.Time
	ProjectID  string
}

type FileType string

const (
	FileTypeNone          FileType = "NONE"
	FileTypeVisualization FileType = "VISUALIZATION"
	FileTypeAlbum         FileType = "ALBUM"
)

func NewFile(
	id, name,
	url string,
	fileType FileType,
	mimeType string,
	size int64,
	projectID string,
) *File {
	var (
		now = time.Now()
	)

	return &File{
		ID:         id,
		Name:       name,
		URL:        url,
		Type:       fileType,
		MimeType:   mimeType,
		Size:       size,
		CreatedAt:  now,
		ModifiedAt: now,
		ProjectID:  projectID,
	}
}
