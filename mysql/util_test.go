package mysql

import (
	"encoding/hex"
	"github.com/flike/kingshard/core/hack"
	"testing"
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
