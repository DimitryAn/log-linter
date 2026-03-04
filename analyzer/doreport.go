package analyzer

import (
	"go/ast"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

func makeReport(bl *ast.BasicLit, p *analysis.Pass, newStr string, messgErr string) {
	p.Report(analysis.Diagnostic{
		Pos:     bl.ValuePos,
		End:     bl.End(),
		Message: messgErr,
		SuggestedFixes: []analysis.SuggestedFix{
			{
				TextEdits: []analysis.TextEdit{
					{
						Pos:     bl.ValuePos,
						End:     bl.End(),
						NewText: []byte(newStr),
					},
				},
			},
		},
	})
}

func makeFix(s string) string {
	var b strings.Builder

	newStr := removeEmoji(s)

	if len(newStr) == 0 {
		return `""`
	}

	for _, v := range newStr {

		_, ok := availableSymbols[v]

		if ok || unicode.Is(unicode.Latin, v) || unicode.IsDigit(v) || unicode.IsSpace(v) {
			b.WriteRune(v)
		}
	}
	newStr = b.String()
	if len(newStr) == 0 {
		return `""`
	}
	return strconv.Quote(strings.ToLower(newStr[:1]) + newStr[1:])
}
