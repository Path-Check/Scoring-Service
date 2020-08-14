package persistence

import (
	"logger/model"
	"os"
	"reflect"
	"testing"
)

func TestOpenFile(t *testing.T) {
	tests := []struct {
		name    string
		want    *os.File
		wantErr bool
	}{
		// TODO: Add test cases.
		// {name: "testfile", want: }
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OpenFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveRequestToFile(t *testing.T) {
	type args struct {
		f   *os.File
		req model.LogRequest
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SaveRequestToFile(tt.args.f, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveRequestToFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SaveRequestToFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloseFile(t *testing.T) {
	type args struct {
		f *os.File
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CloseFile(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("CloseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CloseFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
