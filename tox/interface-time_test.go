package tox

import (
	"reflect"
	"testing"
	"time"

	"gitlab.dohome.technology/dohome-2020/go-servicex/timex"
)

func TestTime(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{name: `1`, args: args{s: "2023-03-16 21:40:30.201622"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Time(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimePOST(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *time.Time
	}{
		{name: `1`, args: args{s: time.Now().Local().Format(timex.MSSQL)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimePOST(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimePOST() = %v, want %v", got, tt.want)
			}
		})
	}
}
