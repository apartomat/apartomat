package files

import "time"

type File struct {
	ID         string
	Name       string
	URL        string
	Type       FileType
	MimeType   string
	CreatedAt  time.Time
	ModifiedAt time.Time
	ProjectID  string
}

type FileType string

const (
	FileTypeNone          FileType = "NONE"
	FileTypeVisualization FileType = "VISUALIZATION"
)

func NewFile(id, name, url string, fileType FileType, mimeType, projectID string) *File {
	now := time.Now()

	return &File{
		ID:         id,
		Name:       name,
		URL:        url,
		Type:       fileType,
		MimeType:   mimeType,
		CreatedAt:  now,
		ModifiedAt: now,
		ProjectID:  projectID,
	}
}
