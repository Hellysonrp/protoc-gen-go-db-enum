// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The protoc-gen-go-db-enum binary is a protoc plugin to generate Go code for
// both proto2 and proto3 versions of the protocol buffer language.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	gengo "github.com/Hellysonrp/protoc-gen-go-db-enum/internal_gengo"
	"google.golang.org/protobuf/compiler/protogen"
)

const version = "1.0.0"

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Fprintf(os.Stdout, "%v %v\n", filepath.Base(os.Args[0]), version)
		os.Exit(0)
	}

	var (
		flags flag.FlagSet
	)
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if f.Generate {
				gengo.GenerateFile(gen, f)
			}
		}
		gen.SupportedFeatures = gengo.SupportedFeatures
		return nil
	})
}
