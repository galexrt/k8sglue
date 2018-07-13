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

package tar

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// CreateTarGz create a ".tar.gz" from the given source path
func CreateTarGz(src string, outputs ...io.Writer) error {
	multiWriter := io.MultiWriter(outputs...)

	gzipWriter := gzip.NewWriter(multiWriter)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	// walk path
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		// return on any filepath.Walk() error
		if err != nil {
			return err
		}

		// create a new directory or file header
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		// update name to reflect the desired destination when untaring
		header.Name = strings.TrimPrefix(strings.Replace(path, src, "", -1), string(filepath.Separator))

		// write the tar header
		if err = tarWriter.WriteHeader(header); err != nil {
			return err
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		// open the file
		f, err := os.Open(path)
		defer f.Close()
		if err != nil {
			return err
		}

		// copy file data into tar writer
		if _, err := io.Copy(tarWriter, f); err != nil {
			return err
		}

		return nil
	})
}
