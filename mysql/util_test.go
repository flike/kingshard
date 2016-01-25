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

package mysql

import (
	"encoding/hex"
	"testing"

	"github.com/flike/kingshard/core/hack"
)

func TestCalcPassword(t *testing.T) {
	/*
		// **** JDBC ****
		seed:
			@jx=d_3z42;sS$YrS)p|
		hex:
			406a783d645f337a34323b73532459725329707c
		pass:
			kingshard
		scramble:
			fbc71db5ac3d7b51048d1a1d88c1677f34bcca11
	*/
	test, _ := RandomBuf(20)
	hex_test := hex.EncodeToString(test)
	t.Logf("rnd seed: %s, %s", hack.String(test), hex_test)

	seed := hack.Slice("@jx=d_3z42;sS$YrS)p|")
	hex_seed := hex.EncodeToString(seed)

	t.Logf("seed: %s equal %s, pass: %v", "406a783d645f337a34323b73532459725329707c", hex_seed, "406a783d645f337a34323b73532459725329707c" == hex_seed)
	scramble := CalcPassword(seed, hack.Slice("kingshard"))

	hex_scramble := hex.EncodeToString(scramble)
	t.Logf("scramble: %s equal %s, pass: %v", "fbc71db5ac3d7b51048d1a1d88c1677f34bcca11", hex_scramble, "fbc71db5ac3d7b51048d1a1d88c1677f34bcca11" == hex_scramble)
}
