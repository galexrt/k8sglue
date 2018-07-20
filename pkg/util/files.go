/*
Copyright (c) 2018 Alexander Trost <galexrt@googlemail.com>. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"os"
	"path/filepath"
	"strconv"
)

// CreateDirectory create a directory if it does not exist
func CreateDirectory(path string, mode string) error {
	n, err := strconv.ParseUint(mode, 8, 32)
	if err != nil {
		return err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.Mkdir(path, os.FileMode(n)); err != nil {
			return err
		}
	}
	return nil
}

// ReturnFullPath return full absolute path for a given path
func ReturnFullPath(path string) (string, error) {
	return filepath.Abs(path)
}

// Symlink create a symlink if it does not exist
func Symlink(oldname, newname string) error {
	if _, err := os.Stat(oldname); os.IsNotExist(err) {
		return err
	}
	if _, err := os.Stat(newname); os.IsNotExist(err) {
		if err = os.Symlink(oldname, newname); err != nil {
			return err
		}
	}
	return nil
}
