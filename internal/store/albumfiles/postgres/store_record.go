package postgres

import (
	. "github.com/apartomat/apartomat/internal/store/albumfiles"
	"github.com/uptrace/bun"
	"time"
)

type record struct {
	bun.BaseModel `bun:"table:apartomat.album_files,alias:af"`

	ID                  string     `bun:"id,pk"`
	Status              string     `bun:"status"`
	AlbumID             string     `bun:"album_id"`
	Version             int        `bun:"version"`
	FileID              *string    `bun:"file_id"`
	GeneratingStartedAt *time.Time `bun:"generating_started_at"`
	GeneratingDoneAt    *time.Time `bun:"generating_done_at"`
	CreatedAt           time.Time  `bun:"created_at"`
	ModifiedAt          time.Time  `bun:"modified_at"`
}

func toRecord(file *AlbumFile) record {
	return record{
		ID:                  file.ID,
		Status:              string(file.Status),
		AlbumID:             file.AlbumID,
		Version:             file.Version,
		FileID:              file.FileID,
		GeneratingStartedAt: file.GeneratingStartedAt,
		GeneratingDoneAt:    file.GeneratingDoneAt,
		CreatedAt:           file.CreatedAt,
		ModifiedAt:          file.ModifiedAt,
	}
}

func toRecords(files []*AlbumFile) []record {
	var (
		records = make([]record, len(files))
	)

	for i, p := range files {
		records[i] = toRecord(p)
	}

	return records
}

func fromRecords(records []record) []*AlbumFile {
	var (
		files = make([]*AlbumFile, len(records))
	)

	for i, rec := range records {
		files[i] = &AlbumFile{
			ID:                  rec.ID,
			Status:              Status(rec.Status),
			AlbumID:             rec.AlbumID,
			Version:             rec.Version,
			FileID:              rec.FileID,
			GeneratingStartedAt: rec.GeneratingStartedAt,
			GeneratingDoneAt:    rec.GeneratingDoneAt,
			CreatedAt:           rec.CreatedAt,
			ModifiedAt:          rec.ModifiedAt,
		}
	}

	return files
}
