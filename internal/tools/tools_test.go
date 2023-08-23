package tools

import (	
	"testing"
)

func TestGetBase(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{struct {
		name string
		args args
		want string
	}{name: "1", args: args{path: "./base/linux.log"}, want: "linux.log"},
		{name: "2", args: args{path: "/base/windows.log"}, want: "windows.log"},
		{name: "3", args: args{path: "/base/windows"}, want: "windows"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBase(tt.args.path); got != tt.want {
				t.Errorf("GetBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileNamewithoutExt(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{struct {
		name string
		args args
		want string
	}{name: "1", args: args{path: "./base/linux.log"}, want: "linux"},
		{name: "2", args: args{path: "/base/windows.log"}, want: "windows"},
		{name: "3", args: args{path: "/base/windows"}, want: "windows"},
	} // TODO: Add test cases.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileNamewithoutExt(tt.args.path); got != tt.want {
				t.Errorf("FileNamewithoutExt() = %v, want %v", got, tt.want)
			} else {
				t.Log(got)
			}
		})
	}
}

func TestGetIntDef(t *testing.T) {
	type args struct {
		value string
		def   int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{struct {
		name string
		args args
		want int64
	}{name: "1", args: args{value: "10", def: 20}, want: 10},
		struct {
			name string
			args args
			want int64
		}{name: "2", args: args{value: "10dsds", def: 20}, want: 20},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIntDef(tt.args.value, tt.args.def); got != tt.want {
				t.Errorf("GetIntDef() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileConfig(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{struct {
		name string
		want string
	}{"1", "config.json"}} // TODO: Add test cases.

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileConfig(); got != tt.want {
				t.Errorf("FileConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArg_String(t *testing.T) {
	tests := []struct {
		name string
		a    Arg
		want string
	}{struct {
		name string
		a    Arg
		want string
	}{name: "1", a: Arg{Fileconfig: "test.config", Debug: true, Help: true}, want: "-config test.config -help -debug"}} // TODO: Add test cases.

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.String(); got != tt.want {
				t.Errorf("Arg.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseFile(t *testing.T) {	

	f:= func(t *testing.T) {
		got, err := ParseFile()
		if err != nil {
			t.Fatal(err)
			return
		}
			t.Log( got)
	}
		t.Run("1",f	)		
					
	
}
