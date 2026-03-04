package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

// изменения регистра буквы
func TestLowercaseAnalyzer(t *testing.T) {
	analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), Analyzer, "lowercase")
}

// отслеживаем не латинскую буквы
func TestOnlyEnglishAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "onlyenglish")
}

// удаление эмоджи
func TestNoEmojiAndSymbols(t *testing.T) {
	analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), Analyzer, "withouttrash")
}

// нахождение чусвтвительных переменных
func TestSensetiveData(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "sensetive")
}

// удаление спец символов
func TestRemoveNonEnglish(t *testing.T) {
	analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), Analyzer, "nonenglish")
}

// особые случаи с пробелами
func TestSpaceBegining(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "space")
}

// замена имени пакета
func TestAliasForPackage(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "alias")
}
