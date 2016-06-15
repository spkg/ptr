package ptr_test

// Do not modify. Generated by nullable-generate.

import (
	"fmt"
	"testing"
	{{range .Imports -}}
	"{{.}}"
	{{- end}}
	"github.com/spkg/ptr"
	"github.com/spkg/ptr/deref"
)

func Test{{.Method}}(t *testing.T) {
	tests := []struct {
		Value {{.NativeType}}
	}{
		{{if eq .NativeType "string" -}}
		{
			Value:     "string-val",
		},
		{{else if eq .NativeType "bool" -}}
		{
			Value:     true,
		},
		{{else if eq .NativeType "time.Time" -}}
		{
			Value:     time.Date(2001, 11, 10, 15, 04, 05, 0, time.FixedZone("AEST", 10*3600)),
		},
		{{else -}}
		{
			Value:     11,
		},
		{{end -}}
		{
			Value: {{.ZeroVal}},
		},
	}
	for i, tt := range tests {
		ttName := fmt.Sprintf("test %d", i)
		p := ptr.{{.Method}}(tt.Value)
		v := deref.{{.Method}}(p)
		{{if eq .NativeType "time.Time" -}}
		if !tt.Value.Equal(v) {
			t.Errorf("%s: expected=%v, actual=%v", ttName, tt.Value, v)
		}
		{{ else -}}
		if tt.Value != v {
			t.Errorf("%s: expected=%v, actual=%v", ttName, tt.Value, v)
		}
		{{end -}}
	}
}
