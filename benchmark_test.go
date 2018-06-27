package main

import (
	"testing"
)

func BenchmarkLoadKeywords(b *testing.B) {
	s := newStatistics()

	for n := 0; n < b.N; n++ {
		s.loadKeywords("./test_files/keywords.txt")
	}
}

func BenchmarkLineStats(b *testing.B) {
	s := newStatistics()

	for n := 0; n < b.N; n++ {
		s.lineStats("the quick brown fox jumps over the lazy dog")
	}
}

func BenchmarkReadFiles(b *testing.B) {
	b.Run("One File - Small", func(b *testing.B) {
		s := newStatistics()

		for n := 0; n < b.N; n++ {
			s.readFiles("./test_files/files/file0.txt")
		}
	})

	b.Run("One File - Large", func(b *testing.B) {
		s := newStatistics()

		for n := 0; n < b.N; n++ {
			s.readFiles("./test_files/files/file4.txt")
		}
	})

	b.Run("Multiple Files - Various Sizes", func(b *testing.B) {
		s := newStatistics()

		for n := 0; n < b.N; n++ {
			s.readFiles("./test_files/files/*")
		}
	})
}
