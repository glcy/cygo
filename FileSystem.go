// Copyright 2016 The cy Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cy

import (
	"os"
)

// fileSystem is just like a namespace
type fileSystem struct {
}

// FileSystem is just like a namespace
var FileSystem fileSystem

// IsExist returns a boolen indicating whether the file is alread exist
func (*fileSystem) IsExist(path string) bool {
	exist := false
	f, err := os.Open(path)
	if err != nil && os.IsExist(err) {
		exist = true
	}
	f.Close()
	return exist
}

// RemoveFileSpec removes the trailing file name and backslash from a path, if they are present
// func (*FileSystem) RemoveFileSpec(path string) string {

// }
