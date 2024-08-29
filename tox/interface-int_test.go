package tox

import "testing"

func TestInt(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: `test1`, args: args{s: 200.19}, want: int(200)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.s); got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
