package functionalgo

import (
	"reflect"
	"testing"
)

func TestTuple(t *testing.T) {
	tuple := Tuple[int, string]{fst: 1, snd: "hello"}

	if tuple.fst != 1 {
		t.Errorf("Expected first value to be 1, got %v", tuple.fst)
	}

	if tuple.snd != "hello" {
		t.Errorf("Expected second value to be 'hello', got %v", tuple.snd)
	}
}

func TestFst(t *testing.T) {
	tuple := Tuple[int, string]{fst: 1, snd: "hello"}
	if Fst(tuple) != 1 {
		t.Errorf("Expected Fst to return 1, got %v", Fst(tuple))
	}
}

func TestSnd(t *testing.T) {
	tuple := Tuple[int, string]{fst: 1, snd: "hello"}
	if Snd(tuple) != "hello" {
		t.Errorf("Expected Snd to return 'hello', got %v", Snd(tuple))
	}
}

func TestTake(t *testing.T) {
	type testcase struct {
		name     string
		input    []any
		num      int
		expected []any
	}

	tests := []testcase{
		{
			name:     "take fewer than length",
			input:    []any{1, 2, 3, 4, 5},
			num:      3,
			expected: []any{1, 2, 3},
		},
		{
			name:     "take more than length",
			input:    []any{1, 2, 3},
			num:      5,
			expected: []any{1, 2, 3},
		},
		{
			name:     "take zero elements",
			input:    []any{1, 2, 3},
			num:      0,
			expected: []any{},
		},
		{
			name:     "take from empty slice",
			input:    []any{},
			num:      2,
			expected: []any{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Take(tc.input, tc.num)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestDrop(t *testing.T) {
	type testcase struct {
		name     string
		input    []any
		num      int
		expected []any
	}

	tests := []testcase{
		{
			name:     "drop fewer than length",
			input:    []any{1, 2, 3, 4, 5},
			num:      2,
			expected: []any{4, 5},
		},
		{
			name:     "drop more than length",
			input:    []any{1, 2, 3},
			num:      5,
			expected: []any{},
		},
		{
			name:     "drop zero elements",
			input:    []any{1, 2, 3},
			num:      0,
			expected: []any{2, 3},
		},
		{
			name:     "drop from empty slice",
			input:    []any{},
			num:      2,
			expected: []any{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Drop(tc.input, tc.num)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestHead(t *testing.T) {
	t.Run("head of non-empty slice", func(t *testing.T) {
		result := Head([]int{1, 2, 3})
		if result != 1 {
			t.Errorf("Expected 1, got %v", result)
		}
	})

	t.Run("head of single element slice", func(t *testing.T) {
		result := Head([]string{"hello"})
		if result != "hello" {
			t.Errorf("Expected 'hello', got %v", result)
		}
	})

	t.Run("head of empty slice panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected Head of empty slice to panic")
			}
		}()
		Head([]int{})
	})
}

func TestTail(t *testing.T) {
	t.Run("tail of non-empty slice", func(t *testing.T) {
		result := Tail([]int{1, 2, 3})
		expected := []int{2, 3}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("tail of single element slice", func(t *testing.T) {
		result := Tail([]string{"hello"})
		expected := []string{}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("tail of empty slice panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected Tail of empty slice to panic")
			}
		}()
		Tail([]int{})
	})
}

func TestMap(t *testing.T) {
	t.Run("map over integers", func(t *testing.T) {
		result := Map(func(x int) int { return x * 2 }, []int{1, 2, 3})
		expected := []int{2, 4, 6}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("map with type conversion", func(t *testing.T) {
		result := Map(func(x int) string { return string([]byte{byte(x + 64)}) }, []int{1, 2, 3})
		expected := []string{"A", "B", "C"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("map over empty slice", func(t *testing.T) {
		result := Map(func(x int) int { return x * 2 }, []int{})
		expected := []int{}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestFilter(t *testing.T) {
	t.Run("filter even numbers", func(t *testing.T) {
		result := Filter(func(x int) bool { return x%2 == 0 }, []int{1, 2, 3, 4, 5, 6})
		expected := []int{2, 4, 6}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("filter with no matches", func(t *testing.T) {
		result := Filter(func(x int) bool { return x > 10 }, []int{1, 2, 3})
		expected := []int{}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("filter all elements match", func(t *testing.T) {
		result := Filter(func(x int) bool { return x < 10 }, []int{1, 2, 3})
		expected := []int{1, 2, 3}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("filter empty slice", func(t *testing.T) {
		result := Filter(func(x int) bool { return x%2 == 0 }, []int{})
		expected := []int{}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestCompose(t *testing.T) {
	t.Run("compose two functions", func(t *testing.T) {
		f := func(x int) float64 { return float64(x) * 1.5 }
		g := func(x float64) string { return string([]byte{byte(int(x) + 64)}) }
		composed := Compose(g, f)

		result := composed(4)
		expected := "F" // 4 * 1.5 = 6 + 64 = 70 = 'F'

		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestFoldl(t *testing.T) {
	t.Run("sum of integers", func(t *testing.T) {
		result := Foldl(func(acc int, x int) int { return acc + x }, 0, []int{1, 2, 3, 4, 5})
		expected := 15
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("string concatenation", func(t *testing.T) {
		result := Foldl(func(acc string, x string) string { return acc + x }, "", []string{"a", "b", "c"})
		expected := "abc"
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("empty slice returns accumulator", func(t *testing.T) {
		result := Foldl(func(acc int, x int) int { return acc + x }, 10, []int{})
		expected := 10
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestFoldr(t *testing.T) {
	t.Run("right-associative concatenation", func(t *testing.T) {
		result := Foldr(
			func(x string, acc string) string { return x + acc },
			"",
			[]string{"a", "b", "c"},
		)
		expected := "abc"
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("foldr for division", func(t *testing.T) {
		// (1 / (2 / (3 / 1))) = 1 / (2 / 3) = 1 / 0.6666 = 1.5
		result := Foldr(
			func(x float64, acc float64) float64 { return x / acc },
			1.0,
			[]float64{1, 2, 3},
		)
		expected := 1.5
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("single element slice", func(t *testing.T) {
		result := Foldr(func(x int, acc int) int { return x + acc }, 5, []int{10})
		expected := 15
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestZipWith(t *testing.T) {
	t.Run("add two slices", func(t *testing.T) {
		result := ZipWith(
			func(a int, b int) int { return a + b },
			[]int{1, 2, 3},
			[]int{4, 5, 6},
		)
		expected := []int{5, 7, 9}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("zip with different types", func(t *testing.T) {
		result := ZipWith(
			func(a int, b string) string { return string([]byte{byte(a + 64)}) + b },
			[]int{1, 2, 3},
			[]string{"x", "y", "z"},
		)
		expected := []string{"Ax", "By", "Cz"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("zip with different length slices", func(t *testing.T) {
		result := ZipWith(
			func(a int, b int) int { return a + b },
			[]int{1, 2, 3, 4},
			[]int{5, 6},
		)
		expected := []int{6, 8}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("zip with empty slice", func(t *testing.T) {
		result := ZipWith(
			func(a int, b int) int { return a + b },
			[]int{1, 2, 3},
			[]int{},
		)
		var expected []int = nil
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestZip(t *testing.T) {
	t.Run("zip two slices", func(t *testing.T) {
		result := Zip([]int{1, 2, 3}, []string{"a", "b", "c"})
		expected := []Tuple[int, string]{
			{fst: 1, snd: "a"},
			{fst: 2, snd: "b"},
			{fst: 3, snd: "c"},
		}

		if len(result) != len(expected) {
			t.Errorf("Expected length %d, got %d", len(expected), len(result))
			return
		}

		for i := range result {
			if result[i].fst != expected[i].fst || result[i].snd != expected[i].snd {
				t.Errorf("At index %d, expected %v, got %v", i, expected[i], result[i])
			}
		}
	})

	t.Run("zip with different length slices", func(t *testing.T) {
		result := Zip([]int{1, 2, 3}, []string{"a", "b"})
		expected := []Tuple[int, string]{
			{fst: 1, snd: "a"},
			{fst: 2, snd: "b"},
		}

		if len(result) != len(expected) {
			t.Errorf("Expected length %d, got %d", len(expected), len(result))
			return
		}

		for i := range result {
			if result[i].fst != expected[i].fst || result[i].snd != expected[i].snd {
				t.Errorf("At index %d, expected %v, got %v", i, expected[i], result[i])
			}
		}
	})

	t.Run("zip with empty slice", func(t *testing.T) {
		result := Zip([]int{1, 2, 3}, []string{})
		var expected []Tuple[int, string] = nil
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestLast(t *testing.T) {
	t.Run("last of non-empty slice", func(t *testing.T) {
		result := Last([]int{1, 2, 3})
		if result != 3 {
			t.Errorf("Expected 3, got %v", result)
		}
	})

	t.Run("last of single element slice", func(t *testing.T) {
		result := Last([]string{"hello"})
		if result != "hello" {
			t.Errorf("Expected 'hello', got %v", result)
		}
	})

	t.Run("last with different types", func(t *testing.T) {
		result := Last([]float64{1.1, 2.2, 3.3})
		expected := 3.3
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("last of empty slice panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic when calling Last on empty slice")
			}
		}()
		Last([]int{})
	})
}
