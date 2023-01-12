package asg

import (
	"testing"

	"github.com/kaisawind/cobol/format"
	"github.com/kaisawind/cobol/options"
)

func TestAnalyzeFile(t *testing.T) {
	opts := options.NewOptions().SetFormat(format.TANDEM)
	program := AnalyzeFile("./testdata/HelloWorld.cbl", opts)
	t.Log(program)
}
