package document

import (
	"regexp"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kaisawind/cobol/gen/preprocessor"
)

type Store struct {
	replaceable preprocessor.IReplaceableContext
	replacement preprocessor.IReplacementContext
}

func NewStore(replaceable preprocessor.IReplaceableContext, replacement preprocessor.IReplacementContext) *Store {
	return &Store{
		replaceable: replaceable,
		replacement: replacement,
	}
}

func (s *Store) GetRegex(replaceable string) (ret string) {
	if replaceable != "" {
		parts := strings.Fields(replaceable)
		regexParts := []string{}
		for _, v := range parts {
			regexParts = append(regexParts, "\\Q"+v+"\\E")
		}
		sep := "[\\r\\n\\s]+"
		ret = strings.Join(regexParts, sep)
	}
	return
}

func (s *Store) Replace(new string, cts *antlr.CommonTokenStream) (ret string) {
	replaceableText := s.getReplaceableText(s.replaceable, cts)
	replacementText := s.getReplacementText(s.replacement, cts)
	if replaceableText != "" && replacementText != "" {
		regexString := s.GetRegex(replaceableText)
		ret = regexp.MustCompile(regexString).ReplaceAllString(new, replacementText)
	} else {
		ret = new
	}
	return
}

func (s *Store) getReplaceableText(ictx preprocessor.IReplaceableContext, cts *antlr.CommonTokenStream) (ret string) {
	ctx, ok := ictx.(*preprocessor.ReplaceableContext)
	if ok {
		if cctx := ctx.PseudoText(); cctx != nil {
			ret = s.extractPseudoText(cctx, cts)
		} else if cctx := ctx.CharDataLine(); cctx != nil {
			ret = GetTextWithHiddenTokens(cctx, cts)
		} else if cctx := ctx.CobolWord(); cctx != nil {
			ret = cctx.GetText()
		} else if cctx := ctx.Literal(); cctx != nil {
			ret = cctx.GetText()
		}
	}
	return
}

func (s *Store) getReplacementText(ictx preprocessor.IReplacementContext, cts *antlr.CommonTokenStream) (ret string) {
	ctx, ok := ictx.(*preprocessor.ReplacementContext)
	if ok {
		if cctx := ctx.PseudoText(); cctx != nil {
			ret = s.extractPseudoText(cctx, cts)
		} else if cctx := ctx.CharDataLine(); cctx != nil {
			ret = GetTextWithHiddenTokens(cctx, cts)
		} else if cctx := ctx.CobolWord(); cctx != nil {
			ret = cctx.GetText()
		} else if cctx := ctx.Literal(); cctx != nil {
			ret = cctx.GetText()
		}
	}
	return
}

func (s *Store) extractPseudoText(ctx preprocessor.IPseudoTextContext, cts *antlr.CommonTokenStream) string {
	pseudoText := GetTextWithHiddenTokens(ctx, cts)

	pseudoText = regexp.MustCompile("^==").ReplaceAllString(pseudoText, "")
	pseudoText = regexp.MustCompile("==$").ReplaceAllString(pseudoText, "")
	pseudoText = strings.TrimSpace(pseudoText)
	return pseudoText
}

type Stores []*Store

func (ss Stores) Len() int {
	return len(ss)
}

func (ss Stores) Less(i, j int) bool {
	iText := ss[i].replaceable.GetText()
	jText := ss[j].replaceable.GetText()
	return len(iText) < len(jText)
}

func (ss Stores) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
