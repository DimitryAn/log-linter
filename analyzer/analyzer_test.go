package analyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestLowercaseAnalyzer(t *testing.T) {
	analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), Analyzer, "lowercase")
}

func TestOnlyEnglishAnalyzer(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "onlyenglish")
}

func TestNoEmojiAndSymbols(t *testing.T) {
	analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), Analyzer, "withouttrash")
}

func TestSensetiveData(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "sensetive")
}

func TestRemoveNonEnglish(t *testing.T) {
	analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), Analyzer, "nonenglish")
}

func TestSpaceBegining(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "space")
}

func TestAliasForPackage(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "alias")
}
