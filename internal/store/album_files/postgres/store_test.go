//go:build test && integration
// +build test,integration

package postgres

import (
	"context"
	"database/sql"
	. "github.com/apartomat/apartomat/internal/store/album_files"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_store_List(t *testing.T) {
	var (
		db = bun.NewDB(
			sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("POSTGRES_DSN")))),
			pgdialect.New(),
		)
	)

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.WithEnabled(true),
	))

	var (
		store = NewStore(db)

		now = time.Now()

		file = &AlbumFile{
			ID:                  "album_file_1",
			Status:              StatusNew,
			AlbumID:             "3FpbidLxfInUfJTH9vOtv",
			Version:             3,
			FileID:              nil,
			GeneratingStartedAt: nil,
			GeneratingDoneAt:    nil,
			CreatedAt:           now,
			ModifiedAt:          now,
		}

		ctx = context.Background()
	)

	if err := store.Save(ctx, file); err != nil {
		t.Errorf("can't save album file: %s", err)
	}

	if res, err := store.List(ctx, IDIn("album_file_1"), SortVersionAsc, 1, 0); err != nil {
		t.Errorf("can't list album files: %s", err)
	} else if len(res) == 0 {
		t.Error("can't list album files: no files")
	} else if !reflect.DeepEqual(res[0], file) {
		t.Errorf("album file = %v, want %v", res[0], file)
	}

	if err := store.Delete(ctx, file); err != nil {
		t.Errorf("can't delete album file: %s", err)
	}
}
