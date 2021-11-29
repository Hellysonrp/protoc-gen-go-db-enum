// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package internal_gengo is internal to the protobuf module.
package internal_gengo

import (
	"google.golang.org/protobuf/compiler/protogen"

	"google.golang.org/protobuf/types/pluginpb"
)

// SupportedFeatures reports the set of supported protobuf language features.
var SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

// GenerateVersionMarkers specifies whether to generate version markers.
var GenerateVersionMarkers = true

// Standard library dependencies.
const (
	databaseSqlDriverPackage = protogen.GoImportPath("database/sql/driver")
)

// GenerateFile generates the contents of a .pb.go file.
func GenerateFile(gen *protogen.Plugin, file *protogen.File) {
	f := newFileInfo(file)
	if len(f.allEnums) == 0 {
		return
	}

	filename := file.GeneratedFilenamePrefix + ".pb.db.enum.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	g.P("package ", f.GoPackageName)
	g.P()

	for _, enum := range f.allEnums {
		genEnum(g, f, enum)
	}
}

func genEnum(g *protogen.GeneratedFile, f *fileInfo, e *protogen.Enum) {
	driverValue := g.QualifiedGoIdent(databaseSqlDriverPackage.Ident("Value"))
	g.P("func (x ", e.GoIdent, ") Value() (", driverValue, ", error) {")
	g.P("return int32(x), nil")
	g.P("}")
	g.P()

	driverInt32ConvertValue := g.QualifiedGoIdent(databaseSqlDriverPackage.Ident("Int32.ConvertValue"))
	g.P("func (x *", e.GoIdent, ") Scan(src interface{}) error {")
	g.P("if src == nil {")
	g.P("*x = ", e.GoIdent, "(0)")
	g.P("} else {")
	g.P("v, err := ", driverInt32ConvertValue, "(src)")
	g.P("if err != nil {")
	g.P("return err")
	g.P("}")
	g.P("*x = ", e.GoIdent, "(v.(int64))")
	g.P("}")
	g.P("return nil")
	g.P("}")
	g.P()
}
