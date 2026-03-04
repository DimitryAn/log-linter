package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

func lintRun(p *analysis.Pass) (any, error) {

	for _, file := range p.Files {

		ast.Inspect(file, func(n ast.Node) bool {

			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if call.Args == nil {
				return true
			}

			if len(call.Args) == 0 {
				return true
			}

			exp, ok := call.Fun.(*ast.SelectorExpr)
			if ok {
				if !checkLoggerName(exp, p) { // проверка на наличие логгера
					return true
				}
				if !checkMethodName(exp) { //проверка на уровень логирования
					return true
				}

				for _, arg := range call.Args {

					doChecks(arg, p)
				}

			}

			return true
		})
	}
	return nil, nil
}

// проверка пакета логера
func checkLoggerName(exp *ast.SelectorExpr, p *analysis.Pass) bool { // это логгер?

	if exp.Sel == nil {
		return false
	}

	obj := p.TypesInfo.ObjectOf(exp.Sel)
	if obj == nil {
		return false
	}

	pkg := obj.Pkg()
	if pkg == nil {
		return false
	}

	if _, ok := availableLoggers[pkg.Path()]; ok {
		return true
	}

	return false
}

// проверка доступного метода логера
func checkMethodName(exp *ast.SelectorExpr) bool {

	if exp.Sel == nil {
		return false
	}

	if _, ok := availableLevels[exp.Sel.Name]; ok {
		return true
	}

	return false
}

// проверка на чувствительные данные при записи s
func doChecks(arg ast.Expr, p *analysis.Pass) {

	ast.Inspect(arg, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.Ident:
			if _, ok := bannedKeywords[node.Name]; ok {
				p.Reportf(node.NamePos, sensetiveErr)
			}
		case *ast.BasicLit:
			onlyEnglishLetter(node, p)
			beginLowerCase(node, p)

		}

		return true
	})
}
