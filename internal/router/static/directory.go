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
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
)

type directory struct {
	Pattern string
	http.FileSystem
}

func newDirectory(dirname string, file os.FileInfo) *directory {
	pattern := "/" + file.Name()
	if pattern != "/" && pattern[len(pattern)-1] != '/' {
		pattern += "/"
	}
	return &directory{Pattern: pattern, FileSystem: http.Dir(filepath.Join(dirname, file.Name()))}
}

func (f *directory) serve(r chi.Router) {
	r.Get(f.Pattern, http.RedirectHandler(f.Pattern, http.StatusMovedPermanently).ServeHTTP)
	r.Get(f.Pattern+"*", func(w http.ResponseWriter, r *http.Request) {
		ctx := chi.RouteContext(r.Context())
		prefix := strings.TrimSuffix(ctx.RoutePattern(), "/*")
		fs := http.StripPrefix(prefix, http.FileServer(f))
		fs.ServeHTTP(w, r)
	})
}
