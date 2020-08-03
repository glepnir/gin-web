// Copyright 2020 glepnir. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package finder

import (
	"errors"
	"os"
	"path/filepath"
)

func InferRootDir(target string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var infer func(d string) (string, error)
	infer = func(d string) (string, error) {
		if d == "/" {
			return "", errors.New("Please run in the root dir")
		}
		_, err := os.Stat(d + string(os.PathSeparator) + target)
		if err == nil || os.IsExist(err) {
			return d, nil
		}
		return infer(filepath.Dir(d))
	}
	return infer(cwd)
}
