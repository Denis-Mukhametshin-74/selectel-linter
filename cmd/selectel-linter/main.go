package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Denis-Mukhametshin-74/selectel-linter/internal/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	if len(os.Args) == 2 && strings.HasPrefix(os.Args[1], "-V=") {
		fmt.Printf("%s version devel buildID=1\n", os.Args[0])
		os.Exit(0)
	}

	singlechecker.Main(analyzer.Analyzer)
}
