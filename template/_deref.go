package deref

// DO NOT MODIFY. Generated by nullable-generate.

import (
	{{range .Imports -}}
	"{{.}}"
	{{- end}}
)

// {{.Method}} returns the value pointed to by {{.Var}}p.
// If {{.Var}}p is nil, then this function returns {{.ZeroVal}}.
func {{.Method}}({{.Var}}p *{{.NativeType}}) {{.NativeType}} {
  if {{.Var}}p == nil {
    return {{.ZeroVal}}
  }
	return *{{.Var}}p
}