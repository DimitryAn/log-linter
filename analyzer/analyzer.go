package analyzer

import (
	"flag"
	"loglinter/config"
	"sync"

	"golang.org/x/tools/go/analysis"
)

const (
	defaultConfigPass string = "loglint.yaml"
	flagName          string = "config"
)

var configOnce sync.Once

var availableLevels = map[string]struct{}{
	"DebugLevel":   {},
	"InfoLevel":    {},
	"WarnLevel":    {},
	"ErrorLevel":   {},
	"DPanicLevel":  {},
	"PanicLevel":   {},
	"Fatal":        {},
	"Fatalf":       {},
	"Fatalln":      {},
	"Print":        {},
	"Printf":       {},
	"":             {},
	"Panic":        {},
	"Panicf":       {},
	"Panicln":      {},
	"Default":      {},
	"Debug":        {},
	"DebugContext": {},
	"Info":         {},
	"Error":        {},
	"ErrorContext": {},
	"InfoContext":  {},
	"Log":          {},
	"LogAttrs":     {},
	"Warn":         {},
	"WarnContext":  {},
}

var availableLoggers = map[string]struct{}{
	"log":             {},
	"log/slog":        {},
	"go.uber.org/zap": {},
}

var bannedKeywords = map[string]struct{}{
	"password":   {},
	"apiKey":     {},
	"token":      {},
	"email":      {},
	"secret":     {},
	"userName":   {},
	"userEmail":  {},
	"pass":       {},
	"api":        {},
	"user":       {},
	"passwd":     {},
	"key":        {},
	"api_key":    {},
	"auth":       {},
	"credential": {},
	"session":    {},
	"cookie":     {},
}

var availableSymbols = map[rune]struct{}{}

var Analyzer = &analysis.Analyzer{
	Name:  "loglinter",
	Doc:   "check correct logging message format",
	Run:   run,
	Flags: addFlgs(),
}

func addFlgs() flag.FlagSet {
	fs := flag.NewFlagSet("logLinter", flag.ExitOnError)
	fs.String(flagName, "", "путь до файла с конфигом")
	return *fs
}

func run(p *analysis.Pass) (any, error) {

	configOnce.Do(func() {
		confPath := p.Analyzer.Flags.Lookup(flagName)

		if confPath.Value.String() == "" { //не передали
			c, err := config.Load(defaultConfigPass)

			if err == nil { //есть в текущей директории
				availableLevels, availableLoggers, bannedKeywords, availableSymbols = c.Map()
			}

		} else {
			c, err := config.Load(confPath.Value.String())

			if err != nil {
				p.Reportf(p.Files[0].Pos(), "loglinter: failed to load config %v, will use default rules", confPath.Value.String())
			}
			availableLevels, availableLoggers, bannedKeywords, availableSymbols = c.Map()
		}
	})

	return lintRun(p)
}
