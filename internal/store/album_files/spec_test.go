package album_files

import (
	"reflect"
	"testing"
)

func TestIDIn(t *testing.T) {
	type args struct {
		vals      []string
		albumFile *AlbumFile
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"first",
			args{
				vals:      []string{"123"},
				albumFile: &AlbumFile{ID: "123"},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IDIn(tt.args.vals...).Is(tt.args.albumFile); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IDIn() = %v, want %v", got, tt.want)
			}
		})
	}
}
