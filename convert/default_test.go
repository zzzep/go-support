package convert

import "testing"

func TestToFloat(t *testing.T) {
	type args struct {
		a            any
		defaultValue float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "TestToFloat", args: args{a: "1", defaultValue: 0}, want: 1},
		{name: "TestToFloat", args: args{a: 1, defaultValue: 0}, want: 1},
		{name: "TestToFloat", args: args{a: 1.0, defaultValue: 0}, want: 1},
		{name: "TestToFloat", args: args{a: 1.000, defaultValue: 0}, want: 1},
		{name: "TestToFloat", args: args{a: true, defaultValue: 0}, want: 0},
		{name: "TestToFloat", args: args{a: nil, defaultValue: 0}, want: 0},
		{name: "TestToFloat", args: args{a: "false", defaultValue: 0}, want: 0},
		{name: "TestToFloat", args: args{a: "0", defaultValue: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloat(tt.args.a, tt.args.defaultValue); got != tt.want {
				t.Errorf("ToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	type args struct {
		a            any
		defaultValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "TestToInt", args: args{a: "1", defaultValue: 0}, want: 1},
		{name: "TestToInt", args: args{a: 1, defaultValue: 0}, want: 1},
		{name: "TestToInt", args: args{a: 1.0, defaultValue: 0}, want: 1},
		{name: "TestToInt", args: args{a: 1.000, defaultValue: 0}, want: 1},
		{name: "TestToInt", args: args{a: true, defaultValue: 0}, want: 0},
		{name: "TestToInt", args: args{a: false, defaultValue: 0}, want: 0},
		{name: "TestToInt", args: args{a: nil, defaultValue: 0}, want: 0},
		{name: "TestToInt", args: args{a: "", defaultValue: 0}, want: 0},
		{name: "TestToInt", args: args{a: "erro", defaultValue: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.args.a, tt.args.defaultValue); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToString(t *testing.T) {
	type args struct {
		a            any
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "TestToString", args: args{a: "1", defaultValue: ""}, want: "1"},
		{name: "TestToString", args: args{a: 1, defaultValue: ""}, want: "1"},
		{name: "TestToString", args: args{a: 1.0, defaultValue: ""}, want: "1"},
		{name: "TestToString", args: args{a: 1.000, defaultValue: ""}, want: "1"},
		{name: "TestToString", args: args{a: true, defaultValue: ""}, want: "true"},
		{name: "TestToString", args: args{a: false, defaultValue: ""}, want: "false"},
		{name: "TestToString", args: args{a: nil, defaultValue: ""}, want: ""},
		{name: "TestToString", args: args{a: "", defaultValue: ""}, want: ""},
		{name: "TestToString", args: args{a: "erro", defaultValue: ""}, want: "erro"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToString(tt.args.a, tt.args.defaultValue); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
