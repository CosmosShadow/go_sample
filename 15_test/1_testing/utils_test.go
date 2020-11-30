package utils

import "testing"

// 普通的测试
func TestGenShortID(t *testing.T) {
	shortID, err := GenShortID()
	if shortID == "" || err != nil {
		t.Error("GenShortID failed")
	}
}

// 性能测试
func BenchmarkGenShortID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenShortID()
	}
}