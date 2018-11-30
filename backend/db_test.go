package backend

import (
	. "github.com/flike/kingshard/backend"

	"testing"
)

func TestDB_Ping(t *testing.T) {
	db, err := Open("127.0.0.1:3306", "root", "", "kingshard", 1)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.Ping()
}
