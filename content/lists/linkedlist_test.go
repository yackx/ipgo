package lists

import (
	"reflect"
	"testing"
)

func TestLength(t *testing.T) {
	fixtures := []struct {
		ll   *LinkedList
		want int
	}{
		{&LinkedList{}, 0},
		{&LinkedList{&Node{1, nil}}, 1},
		{&LinkedList{&Node{10, &Node{20, nil}}}, 2},
	}

	for _, f := range fixtures {
		actual := f.ll.Length()
		if actual != f.want {
			t.Errorf("Invalid Length(): %v want=%d actual=%d", f.ll, f.want, actual)
		}
	}
}

func TestHeadEmpty(t *testing.T) {
	ll := &LinkedList{}
	value, err := ll.Head()
	if value != 0 || err == nil {
		t.Errorf("invalid. value=%d, err=%v", value, err)
	}
}

func TestHead(t *testing.T) {
	fixtures := []struct {
		ll *LinkedList
		value int
	}{
		{&LinkedList{&Node{1, nil}}, 1},
		{&LinkedList{&Node{0, &Node{1, nil}}}, 0},
	}
	for _, f := range fixtures {
		value, err := f.ll.Head()
		if value != f.value || err != nil {
			t.Errorf("invalid for %v. want=%d, actual=%d, err=%v", f.ll, f.value, value, err)
		}
	}
}

func TestGetAtEmpty(t *testing.T) {
	ll := &LinkedList{}
	value, err := ll.GetAt(0)
	if value != 0 || err == nil {
		t.Errorf("invalid. value=%d, err=%v", value, err)
	}
}

func TestGetAtOutOfRange(t *testing.T) {
	ll := &LinkedList{&Node{0, &Node{1, nil}}}
	value, err := ll.GetAt(2)
	if value != 0 || err == nil {
		t.Errorf("invalid. value=%d, err=%v", value, err)
	}
}

func TestGetAt(t *testing.T) {
	ll := &LinkedList{&Node{0, &Node{1, &Node{2, nil}}}}
	for i := 0; i <= 2; i++ {
		value, err := ll.GetAt(i)
		if value != i || err != nil {
			t.Errorf("invalid for %d. want=%d, actual=%d, err=%v", i, i, value, err)
		}
	}
}

func TestTailEmpty(t *testing.T) {
	ll := &LinkedList{}
	value, err := ll.Tail()
	if value != 0 || err == nil {
		t.Errorf("invalid. value=%d, err=%v", value, err)
	}
}

func TestTail(t *testing.T) {
	fixtures := []struct {
		ll *LinkedList
		value int
	}{
		{&LinkedList{&Node{1, nil}}, 1},
		{&LinkedList{&Node{0, &Node{1, nil}}}, 1},
	}
	for _, f := range fixtures {
		value, err := f.ll.Tail()
		if value != f.value || err != nil {
			t.Errorf("invalid for %v. want=%d, actual=%d, err=%v", f.ll, f.value, value, err)
		}
	}
}

func TestInsertHead(t *testing.T) {
	fixtures := []struct {
		ll, want *LinkedList
	}{
		{&LinkedList{}, &LinkedList{&Node{10, nil}}},
		{
			&LinkedList{&Node{0, nil}},
			&LinkedList{&Node{10, &Node{0, nil}}},
		},
		{
			&LinkedList{&Node{0, &Node{1, nil}}},
			&LinkedList{&Node{10, &Node{0, &Node{1, nil}}}},
		},
	}
	for _, f := range fixtures {
		f.ll.InsertBeforeHead(10)
		if !reflect.DeepEqual(f.ll, f.want) {
			t.Errorf("Error InsertAfterTail(0): want=%v, actual=%v", f.want, f.ll)
		}
	}

}

func TestAddTail(t *testing.T) {
	fixtures := []struct {
		ll, want *LinkedList
	}{
		{&LinkedList{}, &LinkedList{&Node{1, nil}}},
		{
			&LinkedList{&Node{0, nil}},
			&LinkedList{&Node{0, &Node{1, nil}}},
		},
		{
			&LinkedList{&Node{0, &Node{1, nil}}},
			&LinkedList{&Node{0, &Node{1, &Node{1, nil}}}},
		},
	}
	for _, f := range fixtures {
		f.ll.InsertAfterTail(1)
		if !reflect.DeepEqual(f.ll, f.want) {
			t.Errorf("mismatch: want=%v, actual=%v", f.want, f.ll)
		}
	}
}

func TestDeleteHeadEmptyList(t *testing.T) {
	emptyList := &LinkedList{}
	err := emptyList.DeleteHead()
	if err == nil {
		t.Errorf("error expected")
	}
}

func TestDeleteHead(t *testing.T) {
	fixtures := []struct {
		ll, want *LinkedList
	}{
		{&LinkedList{&Node{0, nil}}, &LinkedList{}},
		{
			&LinkedList{&Node{0, &Node{1, nil}}},
			&LinkedList{&Node{1, nil}},
		},
	}
	for _, f := range fixtures {
		err := f.ll.DeleteHead()
		if err != nil {
			t.Errorf("error %v", err)
		}
		if !reflect.DeepEqual(f.ll, f.want) {
			t.Errorf("mismatch: want=%v, actual=%v", f.want, f.ll)
		}
	}
}