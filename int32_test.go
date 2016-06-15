package ptr_test

// Do not modify. Generated by nullable-generate.

import (
	"fmt"
	"testing"

	"github.com/spkg/ptr"
	"github.com/spkg/ptr/deref"
)

func TestInt32(t *testing.T) {
	tests := []struct {
		Value int32
	}{
		{
			Value: 11,
		},
		{
			Value: 0,
		},
	}
	for i, tt := range tests {
		ttName := fmt.Sprintf("test %d", i)
		p := ptr.Int32(tt.Value)
		v := deref.Int32(p)
		if tt.Value != v {
			t.Errorf("%s: expected=%v, actual=%v", ttName, tt.Value, v)
		}
	}
}
