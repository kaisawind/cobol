package parser

import (
	"io/ioutil"
	"path"
	"strings"

	"github.com/kaisawind/cobol/gen/preprocessor"
	"github.com/kaisawind/cobol/options"
)

type CopyBookFinder struct {
}

func NewCopyBookFinder() *CopyBookFinder {
	return &CopyBookFinder{}
}

type CobolWordCopyBookFinder struct {
	CopyBookFinder
}

func NewCobolWordCopyBookFinder() *CobolWordCopyBookFinder {
	return &CobolWordCopyBookFinder{
		CopyBookFinder: CopyBookFinder{},
	}
}

func (cw *CobolWordCopyBookFinder) GetCopyBook(ctx preprocessor.ICobolWordContext, opts *options.Options) string {
	for _, v := range opts.CopyBookFiles {
		if cw.isMatching(ctx, v, opts) {
			return v
		}
	}

	for _, v := range opts.CopyBookDirectories {
		filepath := cw.GetCopyBookFromDirectory(ctx, v, opts)
		if filepath != "" {
			return filepath
		}
	}
	return ""
}

func (cw *CobolWordCopyBookFinder) GetCopyBookFromDirectory(ctx preprocessor.ICobolWordContext, dir string, opts *options.Options) (ret string) {
	infos, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}
	if len(infos) == 0 {
		return
	}
	for _, v := range infos {
		filepath := path.Join(dir, v.Name())
		if cw.isMatching(ctx, filepath, opts) {
			return filepath
		}
	}
	return
}

func (cw *CobolWordCopyBookFinder) isMatching(ctx preprocessor.ICobolWordContext, filepath string, opts *options.Options) bool {
	identifier := ctx.GetText()
	if len(opts.CopyBookExtensions) != 0 {
		for _, v := range opts.CopyBookExtensions {
			if cw.isMatchingWithExtension(filepath, identifier, v) {
				return true
			}
		}
		return false
	} else {
		return cw.isMatchingWithoutExtension(filepath, identifier)
	}
}

func (cw *CobolWordCopyBookFinder) isMatchingWithExtension(filepath, identifier, extension string) bool {
	copyBookName := identifier
	if filepath != "" {
		copyBookName = identifier + "." + extension
	}
	filename := path.Base(filepath)
	return strings.EqualFold(filename, copyBookName)
}

func (cw *CobolWordCopyBookFinder) isMatchingWithoutExtension(filepath, identifier string) bool {
	ext := path.Ext(filepath)
	filename := path.Base(filepath)
	basename := strings.TrimSuffix(filename, ext)
	return strings.EqualFold(basename, identifier)
}
