package increment

import "testing"

func TestIncrement(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"normal case",
			args{10},
			55,
		},
		{
			"zero case",
			args{0},
			0,
		},
		{
			"negative integer",
			args{-5},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Increment(tt.args.n)
			if GlobalVar != tt.want {
				t.Errorf("GlobalVar = %v, want %v", GlobalVar, tt.want)
			}
		})

		GlobalVar = 0
	}
}
