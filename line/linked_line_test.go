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
	source := ll
	for {
		if source == nil {
			break
		}
		t.Log(source)
		source = source.next
	}
	source = ll
	code := CombineLinkedLine(source)
	err = os.WriteFile("./testdata/lbli0420.after", []byte(code), os.ModePerm)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
