package sol

import (
	"reflect"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	for idx := 0; idx < b.N; idx++ {
		threeSum(nums)
	}
}
func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "nums = [-1,0,1,2,-1,-4]",
			args: args{nums: []int{-1, 0, 1, 2, -1, -4}},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "nums = [0,1,1]",
			args: args{nums: []int{0, 1, 1}},
			want: [][]int{},
		},
		{
			name: "nums = [0,0,0]",
			args: args{nums: []int{0, 0, 0}},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "nums = [-2,0,1,1,2]",
			args: args{nums: []int{-2, 0, 1, 1, 2}},
			want: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
