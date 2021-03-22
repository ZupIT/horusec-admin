// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package static

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/go-chi/chi"
)

const relativePath = "web/static"

type Assets []*directory

func ListAssets() (Assets, error) {
	files, dir, err := listFiles()
	if err != nil {
		return nil, err
	}

	assets := make(Assets, 0)
	for _, f := range files {
		if f.IsDir() {
			assets = append(assets, newDirectory(dir, f))
		}
	}

	return assets, nil
}

func (a Assets) Serve(r chi.Router) {
	for _, f := range a {
		f.serve(r)
	}
}

func listFiles() (files []os.FileInfo, dir string, err error) {
	dir, err = os.Getwd()
	if err != nil {
		return nil, dir, fmt.Errorf("failed to list files on static directory: %w", err)
	}

	dir = path.Join(dir, relativePath)
	f, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, dir, fmt.Errorf("failed to list files on static directory: %w", err)
	}

	return f, dir, nil
}
