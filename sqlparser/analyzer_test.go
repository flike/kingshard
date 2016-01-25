// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Copyright 2016 The kingshard Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package sqlparser

import "testing"

func TestGetDBName(t *testing.T) {
	wantYes := []string{
		"insert into a.b values(1)",
		"update a.b set c=1",
		"delete from a.b where c=d",
	}
	for _, stmt := range wantYes {
		result, err := GetDBName(stmt)
		if err != nil {
			t.Errorf("error %v on %s", err, stmt)
			continue
		}
		if result != "a" {
			t.Errorf("want a, got %s", result)
		}
	}

	wantNo := []string{
		"insert into a values(1)",
		"update a set c=1",
		"delete from a where c=d",
	}
	for _, stmt := range wantNo {
		result, err := GetDBName(stmt)
		if err != nil {
			t.Errorf("error %v on %s", err, stmt)
			continue
		}
		if result != "" {
			t.Errorf("want '', got %s", result)
		}
	}

	wantErr := []string{
		"select * from a",
		"syntax error",
	}
	for _, stmt := range wantErr {
		_, err := GetDBName(stmt)
		if err == nil {
			t.Errorf("want error, got nil")
		}
	}
}
