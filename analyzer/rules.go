package analyzer

import (
	"go/ast"
	"go/token"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

const (
	availibleSymbolsErr string = "message must contains only english letters, <space> and digits"
	lowercaseErr        string = "message must start with lowercase letter"
	sensetiveErr        string = "potential sensitive data, please rename variable"
	asciiz                     = 122
)

// вспомогательная функция для проверки
func onlyEnglishLetter(bl *ast.BasicLit, p *analysis.Pass) {
	switch bl.Kind {
	case token.STRING:
		str := strings.Trim(bl.Value, `"`)

		if len(removeEmoji(str)) != len(str) {
			newStr := makeFix(str)
			makeReport(bl, p, newStr, availibleSymbolsErr)
			return
		}

		for _, r := range str {

			_, ok := availableSymbols[r]

			if !isAlphabet(r) && !unicode.IsSpace(r) && !unicode.IsDigit(r) && !ok {
				newStr := makeFix(str)
				makeReport(bl, p, newStr, availibleSymbolsErr)
				return
			}
		}
	case token.CHAR:
		r := rune(bl.Value[1])
		_, ok := availableSymbols[r]

		if !isAlphabet(r) && !unicode.IsSpace(r) && !unicode.IsDigit(r) && !ok {
			p.Reportf(bl.ValuePos, availibleSymbolsErr)
		}
	default:
		p.Reportf(bl.ValuePos, availibleSymbolsErr)
	}
}

func beginLowerCase(bl *ast.BasicLit, p *analysis.Pass) {
	switch bl.Kind {
	case token.STRING:

		str := strings.Trim(bl.Value, `"`)
		if len(str) == 0 {
			return
		}

		r := rune(str[0])
		if unicode.IsLetter(r) && !unicode.IsLower(r) && unicode.Is(unicode.Latin, r) && str[0] <= asciiz {
			newStr := makeFix(str)
			makeReport(bl, p, newStr, lowercaseErr)
			return

		} else if unicode.IsLetter(r) && !unicode.Is(unicode.Latin, r) {
			p.Reportf(bl.ValuePos, availibleSymbolsErr)
		} else if !unicode.IsLetter(r) {
			p.Reportf(bl.ValuePos, lowercaseErr)
		}
	case token.CHAR:
		r := rune(bl.Value[1])
		if unicode.IsLetter(r) && !unicode.IsLower(r) {
			p.Reportf(bl.ValuePos, lowercaseErr)
		}
	default:
		p.Reportf(bl.ValuePos, availibleSymbolsErr)
	}
}

func isAlphabet(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func removeEmoji(s string) string {
	emojiPattern := "[\U0001F600-\U0001F64F\U0001F300-\U0001F5FF\U0001F680-\U0001F6FF\U0001F700-\U0001F77F\U0001F780-\U0001F7FF\U0001F800-\U0001F8FF\U0001F900-\U0001F9FF\U00002702-\U000027B0\U000024C2-\U0001F251]+"
	re := regexp.MustCompile(emojiPattern)
	return re.ReplaceAllString(s, "")
}
