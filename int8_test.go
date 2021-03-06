package ptr_test

// Do not modify. Generated by nullable-generate.

import (
	"fmt"
	"testing"

	"github.com/spkg/ptr"
	"github.com/spkg/ptr/deref"
)

func TestInt8(t *testing.T) {
	tests := []struct {
		Value int8
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
		p := ptr.Int8(tt.Value)
		v := deref.Int8(p)
		if tt.Value != v {
			t.Errorf("%s: expected=%v, actual=%v", ttName, tt.Value, v)
		}
	}
}
