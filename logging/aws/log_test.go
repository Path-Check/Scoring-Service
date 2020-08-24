package main

import (
	"reflect"
	"testing"
)

func TestLogger(t *testing.T) {
	type args struct {
		req *LogRequest
	}
	tests := []struct {
		name    string
		args    args
		want    LogResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Logger(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Logger() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Logger() = %v, want %v", got, tt.want)
			}
		})
	}
}
