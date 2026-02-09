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

func TestToPtr(t *testing.T) {
	t.Run("string to *string", func(t *testing.T) {
		val := "hello"
		ptr := ToPtr(val)
		if ptr == nil {
			t.Error("ToPtr() returned nil pointer")
		}
		if *ptr != "hello" {
			t.Errorf("ToPtr() = %v, want %v", *ptr, "hello")
		}
	})

	t.Run("int to *int", func(t *testing.T) {
		val := 42
		ptr := ToPtr(val)
		if ptr == nil {
			t.Error("ToPtr() returned nil pointer")
		}
		if *ptr != 42 {
			t.Errorf("ToPtr() = %v, want %v", *ptr, 42)
		}
	})

	t.Run("float64 to *float64", func(t *testing.T) {
		val := 3.14
		ptr := ToPtr(val)
		if ptr == nil {
			t.Error("ToPtr() returned nil pointer")
		}
		if *ptr != 3.14 {
			t.Errorf("ToPtr() = %v, want %v", *ptr, 3.14)
		}
	})

	t.Run("bool to *bool", func(t *testing.T) {
		val := true
		ptr := ToPtr(val)
		if ptr == nil {
			t.Error("ToPtr() returned nil pointer")
		}
		if *ptr != true {
			t.Errorf("ToPtr() = %v, want %v", *ptr, true)
		}
	})

	t.Run("zero value int to *int", func(t *testing.T) {
		val := 0
		ptr := ToPtr(val)
		if ptr == nil {
			t.Error("ToPtr() returned nil pointer")
		}
		if *ptr != 0 {
			t.Errorf("ToPtr() = %v, want %v", *ptr, 0)
		}
	})

	t.Run("zero value string to *string", func(t *testing.T) {
		val := ""
		ptr := ToPtr(val)
		if ptr == nil {
			t.Error("ToPtr() returned nil pointer")
		}
		if *ptr != "" {
			t.Errorf("ToPtr() = %v, want %v", *ptr, "")
		}
	})

	t.Run("struct to *struct", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		val := Person{Name: "John", Age: 30}
		ptr := ToPtr(val)
		if ptr == nil {
			t.Error("ToPtr() returned nil pointer")
		}
		if ptr.Name != "John" || ptr.Age != 30 {
			t.Errorf("ToPtr() = %v, want %v", *ptr, val)
		}
	})

	t.Run("slice to *slice", func(t *testing.T) {
		val := []int{1, 2, 3}
		ptr := ToPtr(val)
		if ptr == nil {
			t.Error("ToPtr() returned nil pointer")
		}
		if len(*ptr) != 3 || (*ptr)[0] != 1 {
			t.Errorf("ToPtr() = %v, want %v", *ptr, val)
		}
	})
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
