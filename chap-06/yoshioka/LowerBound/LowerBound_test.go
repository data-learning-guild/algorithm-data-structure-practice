func TestLowerBound(t *testing.T) {
	type args struct {
		array  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test0",
			args: args{array: []int{1, 10, 10, 10, 20, 30, 40, 50, 60}, target: 10},
			want: 1,
		},
		{
			name: "test1",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 1},
			want: 0,
		},
		{
			name: "test2",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 50},
			want: 5,
		},
		{
			name: "test3",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 60},
			want: 6,
		},
		{
			name: "test4",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 61},
			want: 7,
		},
		{
			name: "test5",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 2},
			want: 1,
		},
		{
			name: "test6",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60}, target: 59},
			want: 6,
		},
		{
			name: "test7",
			args: args{array: []int{1, 10, 20, 30, 40, 50, 60, 60, 60}, target: 60},
			want: 6,
		},
		{
			name: "test8",
			args: args{array: []int{}, target: 1},
			want: 0,
		},
		{
			name: "test9",
			args: args{array: []int{-5, -2, -2, -2}, target: -4},
			want: 1,
		},
		{
			name: "test10",
			args: args{array: []int{1, 5, 5, 5, 5, 7, 7, 7, 7, 9}, target: 0},
			want: 0,
		},
		{
			name: "test11",
			args: args{array: []int{1, 2, 5, 8, 10}, target: 11},
			want: 5,
		},
		{
			name: "test12",
			args: args{array: []int{1, 2, 2, 2, 3, 3, 10}, target: 3},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowerBound(tt.args.array, tt.args.target); got != tt.want {
				t.Errorf("LowerBound() = %v, want %v", got, tt.want)
			}
		})
	}
}
