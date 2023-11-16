package test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const numbers = 10000

func TestAdd(t *testing.T) {
	//re := add(1, 3)
	//if re != 4 {
	//	t.Errorf("except %d,actual %d", 4, re)
	//}
	var dataset = []struct {
		a   int
		b   int
		out int
	}{
		{1, 1, 2},
		{12, 12, 24},
		{-9, 8, -1},
		{0, 0, 0},
	}

	for _, value := range dataset {
		re := add(value.a, value.b)
		if re != value.out {
			t.Errorf("except %d,actual %d", value.out, re)
		}
	}
}

func TestAdd2(t *testing.T) {
	if testing.Short() {
		t.Skipf("Short 模式下跳过")
	}
	re := add(1, 5)
	if re != 6 {
		t.Errorf("except %d,actual %d", 6, re)
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < numbers; j++ {
			builder.WriteString(strconv.Itoa(j))

		}
		_ = builder.String()
	}
	b.StopTimer()
}
func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = str + strconv.Itoa(j)
		}
	}
	b.StopTimer()
}
func BenchmarkStringSprintf(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < numbers; j++ {
			str = fmt.Sprintf("%s%d", str, j)
		}
	}
	b.StopTimer()
}
