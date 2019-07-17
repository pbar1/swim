package usas

import (
	"io"
	"testing"
)

func Test_extractJS(t *testing.T) {
	type args struct {
		htmlReader io.Reader
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			extractJS(tt.args.htmlReader)
		})
	}
}
