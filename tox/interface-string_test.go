package tox

import (
	"testing"
	"time"
)

func TestString(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: `1`, args: args{s: time.Now()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.s); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
