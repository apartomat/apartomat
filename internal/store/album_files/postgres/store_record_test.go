//go:build test && unit
// +build test,unit

package postgres

import (
	. "github.com/apartomat/apartomat/internal/store/album_files"
	"reflect"
	"testing"
	"time"
)

func Test_toRecord(t *testing.T) {
	type args struct {
		file *AlbumFile
	}

	tests := []struct {
		name string
		args args
		want record
	}{
		{
			"should map AlbumFile to record",
			args{
				file: &AlbumFile{
					ID:                  "album_file_1",
					Status:              StatusNew,
					AlbumID:             "album_id_1",
					Version:             1,
					FileID:              nil,
					GeneratingStartedAt: nil,
					GeneratingDoneAt:    nil,
					CreatedAt:           time.Unix(1257894000, 0),
					ModifiedAt:          time.Unix(1257894001, 0),
				},
			},
			record{
				ID:                  "album_file_1",
				Status:              "NEW",
				AlbumID:             "album_id_1",
				Version:             1,
				FileID:              nil,
				GeneratingStartedAt: nil,
				GeneratingDoneAt:    nil,
				CreatedAt:           time.Unix(1257894000, 0),
				ModifiedAt:          time.Unix(1257894001, 0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toRecord(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_toRecords(t *testing.T) {
	type args struct {
		files []*AlbumFile
	}
	tests := []struct {
		name string
		args args
		want []record
	}{
		{
			"should map list of AlbumFile to records",
			args{
				[]*AlbumFile{
					&AlbumFile{
						ID:                  "album_file_1",
						Status:              StatusNew,
						AlbumID:             "album_id_1",
						Version:             1,
						FileID:              nil,
						GeneratingStartedAt: nil,
						GeneratingDoneAt:    nil,
						CreatedAt:           time.Unix(1257894000, 0),
						ModifiedAt:          time.Unix(1257894001, 0),
					},
					&AlbumFile{
						ID:                  "album_file_2",
						Status:              StatusNew,
						AlbumID:             "album_id_2",
						Version:             3,
						FileID:              nil,
						GeneratingStartedAt: nil,
						GeneratingDoneAt:    nil,
						CreatedAt:           time.Unix(1257894002, 0),
						ModifiedAt:          time.Unix(1257894003, 0),
					},
				},
			},
			[]record{
				{
					ID:                  "album_file_1",
					Status:              "NEW",
					AlbumID:             "album_id_1",
					Version:             1,
					FileID:              nil,
					GeneratingStartedAt: nil,
					GeneratingDoneAt:    nil,
					CreatedAt:           time.Unix(1257894000, 0),
					ModifiedAt:          time.Unix(1257894001, 0),
				},
				{
					ID:                  "album_file_2",
					Status:              "NEW",
					AlbumID:             "album_id_2",
					Version:             3,
					FileID:              nil,
					GeneratingStartedAt: nil,
					GeneratingDoneAt:    nil,
					CreatedAt:           time.Unix(1257894002, 0),
					ModifiedAt:          time.Unix(1257894003, 0),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toRecords(tt.args.files); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fromRecords(t *testing.T) {
	type args struct {
		records []record
	}
	tests := []struct {
		name string
		args args
		want []*AlbumFile
	}{
		{
			"should map records to list of AlbumFile",
			args{
				[]record{
					{
						ID:                  "album_file_1",
						Status:              "NEW",
						AlbumID:             "album_id_1",
						Version:             1,
						FileID:              nil,
						GeneratingStartedAt: nil,
						GeneratingDoneAt:    nil,
						CreatedAt:           time.Unix(1257894000, 0),
						ModifiedAt:          time.Unix(1257894001, 0),
					},
					{
						ID:                  "album_file_2",
						Status:              "NEW",
						AlbumID:             "album_id_2",
						Version:             3,
						FileID:              nil,
						GeneratingStartedAt: nil,
						GeneratingDoneAt:    nil,
						CreatedAt:           time.Unix(1257894002, 0),
						ModifiedAt:          time.Unix(1257894003, 0),
					},
				},
			},
			[]*AlbumFile{
				&AlbumFile{
					ID:                  "album_file_1",
					Status:              StatusNew,
					AlbumID:             "album_id_1",
					Version:             1,
					FileID:              nil,
					GeneratingStartedAt: nil,
					GeneratingDoneAt:    nil,
					CreatedAt:           time.Unix(1257894000, 0),
					ModifiedAt:          time.Unix(1257894001, 0),
				},
				&AlbumFile{
					ID:                  "album_file_2",
					Status:              StatusNew,
					AlbumID:             "album_id_2",
					Version:             3,
					FileID:              nil,
					GeneratingStartedAt: nil,
					GeneratingDoneAt:    nil,
					CreatedAt:           time.Unix(1257894002, 0),
					ModifiedAt:          time.Unix(1257894003, 0),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fromRecords(tt.args.records); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fromRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}
