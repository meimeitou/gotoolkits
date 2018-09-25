package set

import (
	"reflect"
	"sync"
	"testing"
)

func TestSet_ToList(t *testing.T) {
	type fields struct {
		set     map[interface{}]bool
		RWMutex sync.RWMutex
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		// TODO: Add test cases.
		{
			name: "1",
			fields: fields{
				set: map[interface{}]bool{
					"2": true,
					"1": true,
				},
			},
			want: []interface{}{"2", "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Set{
				set:     tt.fields.set,
				RWMutex: tt.fields.RWMutex,
			}
			if got := s.ToList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.ToList() = %v, want %v", got, tt.want)
			}
		})
	}
}
