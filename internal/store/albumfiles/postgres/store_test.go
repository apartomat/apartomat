//go:build test && integration
// +build test,integration

package postgres

import (
	"context"
	"database/sql"
	. "github.com/apartomat/apartomat/internal/store/albumfiles"
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

		ctx = context.Background()

		usersTableName      = "apartomat.users"
		workspacesTableName = "apartomat.workspaces"
		projectsTableName   = "apartomat.projects"
		albumsTableName     = "apartomat.albums"
		albumFilesTableName = "apartomat.album_files"

		now = time.Now()
	)

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.WithEnabled(true),
	))

	{
		// Setup

		var (
			user = map[string]interface{}{
				"id":           "user_1",
				"email":        "mail@test.org",
				"full_name":    "John Doe",
				"is_active":    true,
				"use_gravatar": false,
			}

			workspace = map[string]interface{}{
				"id":        "workspace_1",
				"name":      "My Work",
				"is_active": true,
				"user_id":   "user_1",
			}

			project = map[string]interface{}{
				"id":           "project_1",
				"name":         "My Project",
				"status":       "IN_PROGRESS",
				"workspace_id": "workspace_1",
			}

			album = map[string]interface{}{
				"id":         "album_1",
				"name":       "Album",
				"project_id": "project_1",
			}

			albumFile = map[string]interface{}{
				"id":          `album_file_1`,
				"status":      "NEW",
				"album_id":    `album_1`,
				"version":     3,
				"created_at":  now,
				"modified_at": now,
			}
		)

		db.NewInsert().Table(usersTableName).Model(&user).Exec(ctx)

		db.NewInsert().Table(workspacesTableName).Model(&workspace).Exec(ctx)

		db.NewInsert().Table(projectsTableName).Model(&project).Exec(ctx)

		db.NewInsert().Table(albumsTableName).Model(&album).Exec(ctx)

		db.NewInsert().Table(albumFilesTableName).Model(&albumFile).Exec(ctx)
	}

	t.Cleanup(func() {
		db.NewDelete().Table(usersTableName).Where(`id = 'user_1'`).Exec(ctx)
		db.NewDelete().Table(workspacesTableName).Where(`id = 'workspace_1'`).Exec(ctx)
		db.NewDelete().Table(projectsTableName).Where(`id = 'project_1'`).Exec(ctx)
		db.NewDelete().Table(albumsTableName).Where(`id = 'album_1'`).Exec(ctx)
		db.NewDelete().Table(albumFilesTableName).Where(`id = 'album_file_1'`).Exec(ctx)
	})

	var (
		store = NewStore(db)

		file = &AlbumFile{
			ID:                  "album_file_1         ",
			Status:              StatusNew,
			AlbumID:             "album_1              ",
			Version:             3,
			FileID:              nil,
			GeneratingStartedAt: nil,
			GeneratingDoneAt:    nil,
			CreatedAt:           now.UTC(),
			ModifiedAt:          now.UTC(),
		}
	)

	if res, err := store.List(ctx, IDIn("album_file_1"), SortVersionAsc, 1, 0); err != nil {
		t.Errorf("can't list album files: %s", err)
	} else if len(res) == 0 {
		t.Error("can't list album files: no files")
	} else if !reflect.DeepEqual(res[0], file) {
		t.Errorf("album file got %v, want %v", res[0], file)
	}
}
