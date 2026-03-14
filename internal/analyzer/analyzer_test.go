package analyzer

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestCorrectData(t *testing.T) {
	testdata, err := filepath.Abs(filepath.Join("..", "..", "testdata"))
	if err != nil {
		t.Fatal(err)
	}

	analysistest.Run(t, testdata, Analyzer, "correct")
}

func TestIncorrectData(t *testing.T) {
	testdata, err := filepath.Abs(filepath.Join("..", "..", "testdata"))
	if err != nil {
		t.Fatal(err)
	}

	analysistest.Run(t, testdata, Analyzer, "incorrect")
}
