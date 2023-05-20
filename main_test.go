package main

import (
	"testing"
)

func BenchmarkFunnel64_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		funnel64_1()
	}
}

func BenchmarkFunnel64_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		funnel64_2()
	}
}

func BenchmarkFunnel32_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		funnel32_1()
	}
}

func BenchmarkFunnel32_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		funnel32_2()
	}
}
func BenchmarkFunnel16_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		funnel16_1()
	}
}

func BenchmarkFunnel16_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		funnel16_2()
	}
}
func BenchmarkFunnel8_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		funnel8_1()
	}
}

func BenchmarkFunnel8_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		funnel8_2()
	}
}
