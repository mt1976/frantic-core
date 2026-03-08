package banking

import (
	"testing"
)

func TestNewISO23897UTI(t *testing.T) {
	type args struct {
		generatingEntity string
	}
	tests := []struct {
		name    string
		args    args
		want    UTI
		wantErr bool
	}{
		// TODO: Add test cases.
		{"TEST", args{"TEST"}, UTI{""}, true},
		{"12345678901234567890", args{"12345678901234567890"}, UTI{"12345678901234567890"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if _, err := NewISO23897UTI(tt.args.generatingEntity); (err != nil) != tt.wantErr {
				t.Errorf("NewISO23897UTI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
