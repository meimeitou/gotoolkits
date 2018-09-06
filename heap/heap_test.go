package heap

import (
	"reflect"
	"testing"
)

func TestIntHeap_Len(t *testing.T) {
	tests := []struct {
		name string
		h    IntHeap
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Len(); got != tt.want {
				t.Errorf("IntHeap.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntHeap_Less(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    IntHeap
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("IntHeap.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntHeap_Swap(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		h    IntHeap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestIntHeap_Init(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		h    *IntHeap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Init(tt.args.s)
		})
	}
}

func TestIntHeap_Push(t *testing.T) {
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		h    *IntHeap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Push(tt.args.x)
		})
	}
}

func TestIntHeap_Pop(t *testing.T) {
	tests := []struct {
		name string
		h    *IntHeap
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntHeap.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
