// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal_gengo

import (
	"google.golang.org/protobuf/compiler/protogen"
)

type fileInfo struct {
	*protogen.File

	allEnums []*protogen.Enum
}

func newFileInfo(file *protogen.File) *fileInfo {
	f := &fileInfo{File: file}

	// Collect all enums, messages, and extensions in "flattened ordering".
	// See filetype.TypeBuilder.
	var walkMessages func([]*protogen.Message, func(*protogen.Message))
	walkMessages = func(messages []*protogen.Message, f func(*protogen.Message)) {
		for _, m := range messages {
			f(m)
			walkMessages(m.Messages, f)
		}
	}
	initEnumInfos := func(enums []*protogen.Enum) {
		f.allEnums = append(f.allEnums, enums...)
	}
	initEnumInfos(f.Enums)
	walkMessages(f.Messages, func(m *protogen.Message) {
		initEnumInfos(m.Enums)
	})

	return f
}
