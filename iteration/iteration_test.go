package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expect %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	repeated := Repeat("a", 10)
	fmt.Println(repeated)
	// Output: aaaaaaaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
