package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 4)
	expected := "aaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}

func BenchmarkStandardRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat("a", 4)
	}
}
