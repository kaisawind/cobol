package test

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/kaisawind/cobol/document"
	"github.com/kaisawind/cobol/format"
	"github.com/kaisawind/cobol/options"
)

var (
	COBOL_EXTS = []string{".cbl", ".cob", ".jcl", ".txt", ""}
)

func TestPreprocessorExecCics(t *testing.T) {
	rootdir := "./testdata/cobol/preprocessor"
	dirname := "fixed"
	parentdir := path.Join(rootdir, dirname)
	filename := "ExecCics.cbl"

	opts := options.NewOptions().AddCopyBookDirectory(parentdir).SetFormat(dir2format(dirname))
	filepath := path.Join(parentdir, filename)
	processed := document.ParseFile(filepath, opts)
	processedBuf, err := os.ReadFile(filepath + ".preprocessed")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if processed != string(processedBuf) {
		fmt.Println(processed)
		fmt.Println(strings.Repeat("-", 78))
		fmt.Println(string(processedBuf))
		t.FailNow()
	}
}

func TestPreprocessorIdentificationDivision(t *testing.T) {
	rootdir := "./testdata/cobol/preprocessor"
	dirname := "fixed"
	parentdir := path.Join(rootdir, dirname)
	filename := "IdentificationDivisionWithPadding.cbl"

	opts := options.NewOptions().AddCopyBookDirectory(parentdir).SetFormat(dir2format(dirname))
	filepath := path.Join(parentdir, filename)
	processed := document.ParseFile(filepath, opts)
	processedBuf, err := os.ReadFile(filepath + ".preprocessed")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if processed != string(processedBuf) {
		fmt.Println(processed)
		fmt.Println(strings.Repeat("-", 78))
		fmt.Println(string(processedBuf))
		t.FailNow()
	}
}

func TestPreprocessorLineContinuation(t *testing.T) {
	rootdir := "./testdata/cobol/preprocessor"
	dirname := "fixed"
	parentdir := path.Join(rootdir, dirname)
	filename := "LineContinuation.cbl"

	opts := options.NewOptions().AddCopyBookDirectory(parentdir).SetFormat(dir2format(dirname))
	filepath := path.Join(parentdir, filename)
	processed := document.ParseFile(filepath, opts)
	processedBuf, err := os.ReadFile(filepath + ".preprocessed")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if processed != string(processedBuf) {
		fmt.Println(processed)
		fmt.Println(strings.Repeat("-", 78))
		fmt.Println(string(processedBuf))
		t.FailNow()
	}
}

func TestPreprocessorTandemExecCics(t *testing.T) {
	rootdir := "./testdata/cobol/preprocessor"
	dirname := "tandem"
	parentdir := path.Join(rootdir, dirname)
	filename := "ExecCics.cbl"

	opts := options.NewOptions().AddCopyBookDirectory(parentdir).SetFormat(dir2format(dirname))
	filepath := path.Join(parentdir, filename)
	processed := document.ParseFile(filepath, opts)
	processedBuf, err := os.ReadFile(filepath + ".preprocessed")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if processed != string(processedBuf) {
		fmt.Println(processed)
		fmt.Println(strings.Repeat("-", 78))
		fmt.Println(string(processedBuf))
		t.FailNow()
	}
}

func TestPreprocessor(tt *testing.T) {
	rootdir := "./testdata/cobol/preprocessor"
	des, err := os.ReadDir(rootdir)
	if err != nil {
		tt.Error(err)
		tt.FailNow()
	}
	for _, dir := range des {
		if !dir.IsDir() {
			continue
		}
		tt.Run(dir.Name(), func(t *testing.T) {
			parentDir := path.Join(rootdir, dir.Name())
			files, err := os.ReadDir(parentDir)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}
			for _, file := range files {
				if file.IsDir() {
					continue
				}
				ext := path.Ext(file.Name())
				t.Log(ext, file.Name())
				isCobol := false
				for _, v := range COBOL_EXTS {
					if strings.EqualFold(v, ext) {
						isCobol = true
						break
					}
				}
				if !isCobol {
					continue
				}
				opts := options.NewOptions().AddCopyBookDirectory(parentDir).SetFormat(dir2format(dir.Name()))
				filename := path.Join(parentDir, file.Name())
				processed := document.ParseFile(filename, opts)
				processedBuf, err := os.ReadFile(filename + ".preprocessed")
				if err != nil {
					t.Error(err)
					t.FailNow()
				}
				if processed != string(processedBuf) {
					fmt.Println(processed)
					fmt.Println(strings.Repeat("-", 78))
					fmt.Println(string(processedBuf))
					t.FailNow()
				}
			}
		})
	}
}

func dir2format(name string) format.Format {
	f := format.FIXED
	switch name {
	case "fixed":
		f = format.FIXED
	case "tandem":
		f = format.TANDEM
	case "variable":
		f = format.VARIABLE
	default:
		f = format.FIXED
	}
	return f
}
