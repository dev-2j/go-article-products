package tox

import "testing"

func TestBool(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: `nil`, args: args{nil}, want: false},
		{name: `1 int`, args: args{1}, want: true},
		{name: `1 int ptr`, args: args{IntPtr(1)}, want: true},
		{name: `0 int`, args: args{0}, want: false},
		{name: `1 string`, args: args{`1`}, want: true},
		{name: `T`, args: args{`T`}, want: true},
		{name: `true`, args: args{`true`}, want: true},
		{name: `True`, args: args{`True`}, want: true},
		{name: `TRUE`, args: args{`TRUE`}, want: true},
		{name: `Y`, args: args{`Y`}, want: true},
		{name: `Yes`, args: args{`Yes`}, want: true},
		{name: `YES`, args: args{`YES`}, want: true},
		{name: `F`, args: args{`F`}, want: false},
		{name: `false`, args: args{`false`}, want: false},
		{name: `False`, args: args{`False`}, want: false},
		{name: `FALSE`, args: args{`FALSE`}, want: false},
		{name: `0`, args: args{`0`}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bool(tt.args.s); got != tt.want {
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}
