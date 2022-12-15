package line

import (
	"os"
	"testing"

	"github.com/kaisawind/cobol/format"
)

func TestLinkedLine(t *testing.T) {
	f, err := os.Open("./testdata/lbli0420.src")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	defer f.Close()
	ll := NewLinkedLine(f, format.FIXED)
	for {
		if ll == nil {
			break
		}
		t.Log(ll)
		ll = ll.next
	}

}
