/*
 * Copyright (c) 2021 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package modfile

import (
	"testing"
)

func TestSplitFname(t *testing.T) {
	type testCase struct {
		fname string
		ext   string
	}
	cases := []testCase{
		{"foo.spx", ".spx"},
		{"foo_yap.gox", "_yap.gox"},
		{"foo.gox", ".gox"},
	}
	for _, c := range cases {
		if ext := ClassExt(c.fname); ext != c.ext {
			t.Fatalf("ClassExt(%s): expect %s, got %s\n", c.fname, c.ext, ext)
		}
	}
}
