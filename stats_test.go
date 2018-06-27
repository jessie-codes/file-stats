package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStatistics(t *testing.T) {
	s := newStatistics()
	assert.IsType(t, &statistics{}, s)
}

func TestLoadKeywords(t *testing.T) {
	s := newStatistics()
	s.loadKeywords("./test_files/keywords.txt")

	expected := map[string]int{
		"far":  0,
		"some": 0,
		"but":  0,
		"and":  0,
		"a":    0,
	}

	assert.Exactly(t, expected, s.keywords)
}

func TestCalc(t *testing.T) {
	s := newStatistics()
	s.lineLengths = []float64{4, 5, 6}
	s.tokenCounts = []float64{7, 8, 9}
	s.calc()

	assert.Equal(t, float64(5), s.medLength)
	assert.Equal(t, float64(8), s.medTokens)
	assert.Equal(t, float64(0.816496580927726), s.stdLength)
	assert.Equal(t, float64(0.816496580927726), s.stdTokens)
}

func TestLineStats(t *testing.T) {
	s := newStatistics()
	s.lineStats("the quick brown fox jumps over the lazy dog")

	assert.Exactly(t, []float64{43}, s.lineLengths)
	assert.Exactly(t, []float64{9}, s.tokenCounts)
}

func TestReadFiles(t *testing.T) {
	s := newStatistics()
	s.keywords = map[string]int{
		"far":  0,
		"some": 0,
		"but":  0,
		"and":  0,
		"a":    0,
	}

	s.readFiles("./test_files/files/file0.txt")

	expected := statistics{
		keywords: map[string]int{
			"far":  2,
			"some": 0,
			"but":  0,
			"and":  2,
			"a":    2,
		},
		lines: map[string]int{
			"far a and": 2,
			"test zero": 1,
		},
	}

	assert.Equal(t, expected.lines["far a and"], s.lines["far a and"])
	assert.Equal(t, expected.keywords["far"], s.keywords["far"])
}
