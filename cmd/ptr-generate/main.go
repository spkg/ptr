package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var debug = debugT(true)

type debugT bool

func (d debugT) Printf(format string, args ...interface{}) {
	if d {
		s := strings.TrimSpace(fmt.Sprintf(format, args...))
		println(s)
	}
}

func main() {
	log.SetFlags(0)
	ptrDir := ptrDir()
	templateDir := filepath.Join(ptrDir, "template")
	debug.Printf("ptr dir: %s", ptrDir)
	debug.Printf("template dir: %s", templateDir)

	files := []struct {
		template string
		output   string
	}{
		{
			template: "_deref.go",
			output:   filepath.Join("deref", "/xx.go"),
		},
		{
			template: "_ptr_test.go",
			output:   "xx_test.go",
		},
		{
			template: "_ptr.go",
			output:   "xx.go",
		},
	}

	for _, ff := range files {
		debug.Printf("ff.template: %s", ff.template)
		templateFile := filepath.Join(templateDir, ff.template)
		_, err := os.Stat(templateFile)
		if err != nil {
			log.Fatalf("cannot find template file %s", templateFile)
		}
		tmpl, err := template.ParseFiles(templateFile)
		if err != nil {
			log.Fatalf("cannot parse %s: %v", ff.template, err)
		}

		for _, p := range params {
			var outBuf bytes.Buffer
			if err = tmpl.Execute(&outBuf, p); err != nil {
				log.Fatalf("template %s failed for %s: %v", ff.template, p.NativeType, err)
			}
			fileBase := p.File
			if fileBase == "" {
				fileBase = p.NativeType
			}
			outFileName := strings.Replace(ff.output, "xx", fileBase, 1)
			outBytes, err := format.Source(outBuf.Bytes())
			if err != nil {
				log.Fatalf("cannot format %s: %v", outFileName, err)
			}

			outFilePath := filepath.Join(ptrDir, outFileName)
			outFile, err := os.Create(outFilePath)
			if err != nil {
				log.Fatalf("cannot create %s: %v", outFilePath, err)
			}
			if _, err = outFile.Write(outBytes); err != nil {
				log.Fatalf("cannot write to %s: %v", outFilePath, err)
			}
			outFile.Close()
			fmt.Println(outFileName)
		}
	}
}

// ptrDir returns the directory for the ptr package.
// All other directories are relative to this.
func ptrDir() string {
	gopath := os.Getenv("GOPATH")
	relfile := filepath.FromSlash("src/github.com/spkg/ptr")
	for _, godir := range filepath.SplitList(gopath) {
		dir := filepath.Join(godir, relfile)
		s, err := os.Stat(dir)
		if err == nil && s.IsDir() {
			return dir
		}
	}
	log.Fatalf("cannot find %s in $GOPATH", relfile)
	panic("not reached")
}

// params contains parameters for the template.
var params = []struct {
	Method     string   // Method name, eg "Int32", "Time"
	NativeType string   // Native type, eg "int32", "time.Time"
	File       string   // Base name of file
	Indefinite string   // Indefinite article: "a" or "an"
	ZeroVal    string   // Zero value for the type, eg "0", "time.Time{}"
	Var        string   // Variable name,eg "n", "t"
	Imports    []string // Any additional imports required
}{
	{
		Method:     "String",
		NativeType: "string",
		Indefinite: "a",
		ZeroVal:    `""`,
		Var:        "s",
	},
	{
		Method:     "Bool",
		NativeType: "bool",
		Indefinite: "a",
		ZeroVal:    "false",
		Var:        "b",
	},
	{
		Method:     "Time",
		NativeType: "time.Time",
		File:       "time",
		Indefinite: "a",
		ZeroVal:    "time.Time{}",
		Var:        "t",
		Imports:    []string{"time"},
	},
	{
		Method:     "Float64",
		NativeType: "float64",
		Indefinite: "a",
		ZeroVal:    "0",
		Var:        "n",
	},
	{
		Method:     "Float32",
		NativeType: "float32",
		Indefinite: "a",
		ZeroVal:    "0",
		Var:        "n",
	},
	{
		Method:     "Int",
		NativeType: "int",
		Indefinite: "an",
		ZeroVal:    "0",
		Var:        "n",
	},
	{
		Method:     "Uint",
		NativeType: "uint",
		Indefinite: "a",
		ZeroVal:    "0",
		Var:        "n",
	},
	{
		Method:     "Int64",
		NativeType: "int64",
		Indefinite: "an",
		ZeroVal:    "0",
		Var:        "n",
	},
	{
		Method:     "Uint64",
		NativeType: "uint64",
		Indefinite: "a",
		ZeroVal:    "0",
		Var:        "n",
	},
	{
		Method:     "Int32",
		NativeType: "int32",
		Indefinite: "an",
		ZeroVal:    "0",
		Var:        "n",
	},
	{
		Method:     "Uint32",
		NativeType: "uint32",
		Indefinite: "a",
		ZeroVal:    "0",
		Var:        "n",
	},
	{
		Method:     "Int16",
		NativeType: "int16",
		Indefinite: "an",
		ZeroVal:    "0",
		Var:        "n",
	},
	{
		Method:     "Uint16",
		NativeType: "uint16",
		Indefinite: "a",
		ZeroVal:    "0",
		Var:        "n",
	},
	{
		Method:     "Int8",
		NativeType: "int8",
		Indefinite: "an",
		ZeroVal:    "0",
		Var:        "n",
	},
	{
		Method:     "Byte",
		NativeType: "byte",
		Indefinite: "a",
		ZeroVal:    "0",
		Var:        "n",
	},
}
