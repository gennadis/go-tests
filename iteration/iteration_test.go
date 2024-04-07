package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("got %q but expected %q", repeated, expected)
	}
}

func TestCounter(t *testing.T) {
	joined := Counter("hello", "l")
	expected := 2

	if joined != expected {
		t.Errorf("got %q but expected %q", joined, expected)
	}

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
