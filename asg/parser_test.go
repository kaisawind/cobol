package asg

import (
	"testing"

	"github.com/kaisawind/cobol/format"
	"github.com/kaisawind/cobol/options"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestAnalyzeFile(t *testing.T) {
	opts := options.NewOptions().SetFormat(format.FIXED)
	program := AnalyzeFile("./testdata/HelloWorld.cbl", opts)
	// t.Log(program)
	buf, err := protojson.Marshal(program)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(string(buf))
}
