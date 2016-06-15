# Pointers to builtin types.

[![GoDoc](https://godoc.org/github.com/spkg/ptr?status.svg)](https://godoc.org/github.com/spkg/ptr)
[![Build Status (Linux)](https://travis-ci.org/spkg/ptr.svg?branch=master)](https://travis-ci.org/spkg/ptr)
[![Build status (Windows)](https://ci.appveyor.com/api/projects/status/5v82q3kwo0iwp31f?svg=true)](https://ci.appveyor.com/project/jjeffery/ptr)
[![Coverage Status](https://coveralls.io/repos/github/spkg/ptr/badge.svg?branch=master)](https://coveralls.io/github/spkg/ptr?branch=master)
[![GoReportCard](http://goreportcard.com/badge/spkg/ptr)](http://goreportcard.com/report/spkg/ptr)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://raw.githubusercontent.com/spkg/ptr/master/LICENSE.md)

Package `ptr` is a very simple package that provides functions to return pointers
to values of builtin types.

The `deref` subdirectory contains a package for safely dereferencing pointers.

These simple packages can be useful when dealing with APIs and datatypes that use
pointers to represent optional parameters.
