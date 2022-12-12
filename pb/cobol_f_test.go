package pb

import (
	"bytes"
	_ "embed"
	"io"
	"testing"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

//go:embed testdata/lbli0420.src
var test1 []byte

func TestSourceUnit_UnmarshalF(t *testing.T) {
	reader := transform.NewReader(bytes.NewReader(test1), japanese.ShiftJIS.NewDecoder())
	buff, err := io.ReadAll(reader)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	unit := &SourceUnit{}
	err = unit.UnmarshalF(string(buff))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	for _, v := range unit.Contents {
		t.Logf("%s", v.Content)
	}
}
