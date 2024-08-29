package tox

import "testing"

func TestFloat(t *testing.T) {
	type args struct {
		s interface{}
	}
	x := 1234
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: `0 string`, args: args{`0`}, want: 0},
		{name: `0 int`, args: args{0}, want: 0},
		{name: `1.234 string`, args: args{`1.234`}, want: 1.234},
		{name: `1.234 num`, args: args{1.234}, want: 1.234},
		{name: `1,234 พัน`, args: args{`1,234`}, want: 0},
		{name: `1234 ptr`, args: args{&x}, want: 1234},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Float(tt.args.s); got != tt.want {
				t.Errorf("Float() = %v, want %v", got, tt.want)
			}
		})
	}
}
