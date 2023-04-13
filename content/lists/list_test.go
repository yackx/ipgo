package list

import (
	"reflect"
	"testing"
)

func pointer(i int) *int {
	p := new(int)
	*p = i
	return p
}

func TestLen(t *testing.T) {
	tests := []struct {
		input List
		want  int
	}{
		{[]int{}, 0},
		{[]int{1, 2, 3}, 3},
	}

	for _, test := range tests {
		got := test.input.Len()
		if got != test.want {
			t.Errorf("Len() %v=>%d, want=%d", test.input, got, test.want)
		}
	}
}

func TestHead(t *testing.T) {
	tests := []struct {
		input List
		want  *int
	}{
		{[]int{}, nil},
		{[]int{1, 2, 3}, pointer(1)},
	}

	for _, test := range tests {
		got := test.input.Head()
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Head() %v=>%v, want=%v", test.input, got, test.want)
		}
	}
}
