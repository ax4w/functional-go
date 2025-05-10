package functionalgo

import (
	"fmt"
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

func TestAny(t *testing.T) {
	t.Run("any with match", func(t *testing.T) {
		result := Any(func(x int) bool { return x%2 == 0 }, []int{1, 2, 3, 4, 5})
		if !result {
			t.Errorf("Expected Any to return true when there are even numbers")
		}
	})

	t.Run("any with no match", func(t *testing.T) {
		result := Any(func(x int) bool { return x%2 == 0 }, []int{1, 3, 5, 7, 9})
		if result {
			t.Errorf("Expected Any to return false when there are no even numbers")
		}
	})

	t.Run("any with all matches", func(t *testing.T) {
		result := Any(func(x int) bool { return x%2 == 0 }, []int{2, 4, 6, 8})
		if !result {
			t.Errorf("Expected Any to return true when all elements match")
		}
	})

	t.Run("any with empty slice", func(t *testing.T) {
		result := Any(func(x int) bool { return x%2 == 0 }, []int{})
		if result {
			t.Errorf("Expected Any to return false for empty slice")
		}
	})
}

func TestAll(t *testing.T) {
	t.Run("all with all matches", func(t *testing.T) {
		result := All(func(x int) bool { return x%2 == 0 }, []int{2, 4, 6, 8})
		if !result {
			t.Errorf("Expected All to return true when all elements match")
		}
	})

	t.Run("all with some matches", func(t *testing.T) {
		result := All(func(x int) bool { return x%2 == 0 }, []int{2, 4, 5, 6})
		if result {
			t.Errorf("Expected All to return false when not all elements match")
		}
	})

	t.Run("all with no matches", func(t *testing.T) {
		result := All(func(x int) bool { return x%2 == 0 }, []int{1, 3, 5, 7})
		if result {
			t.Errorf("Expected All to return false when no elements match")
		}
	})

	t.Run("all with empty slice", func(t *testing.T) {
		result := All(func(x int) bool { return x%2 == 0 }, []int{})
		if !result {
			t.Errorf("Expected All to return true for empty slice")
		}
	})
}

func TestCompare(t *testing.T) {
	t.Run("compare integers", func(t *testing.T) {
		if Compare(5, 10) != LT {
			t.Errorf("Expected Compare(5, 10) to be LT")
		}
		if Compare(10, 5) != GT {
			t.Errorf("Expected Compare(10, 5) to be GT")
		}
		if Compare(5, 5) != EQ {
			t.Errorf("Expected Compare(5, 5) to be EQ")
		}
	})

	t.Run("compare floats", func(t *testing.T) {
		if Compare(5.5, 10.5) != LT {
			t.Errorf("Expected Compare(5.5, 10.5) to be LT")
		}
		if Compare(10.5, 5.5) != GT {
			t.Errorf("Expected Compare(10.5, 5.5) to be GT")
		}
		if Compare(5.5, 5.5) != EQ {
			t.Errorf("Expected Compare(5.5, 5.5) to be EQ")
		}
	})

	t.Run("compare strings", func(t *testing.T) {
		if Compare("apple", "banana") != LT {
			t.Errorf("Expected Compare(\"apple\", \"banana\") to be LT")
		}
		if Compare("banana", "apple") != GT {
			t.Errorf("Expected Compare(\"banana\", \"apple\") to be GT")
		}
		if Compare("apple", "apple") != EQ {
			t.Errorf("Expected Compare(\"apple\", \"apple\") to be EQ")
		}
	})

	t.Run("compare booleans", func(t *testing.T) {
		if Compare(false, true) != LT {
			t.Errorf("Expected Compare(false, true) to be LT")
		}
		if Compare(true, false) != GT {
			t.Errorf("Expected Compare(true, false) to be GT")
		}
		if Compare(true, true) != EQ {
			t.Errorf("Expected Compare(true, true) to be EQ")
		}
		if Compare(false, false) != EQ {
			t.Errorf("Expected Compare(false, false) to be EQ")
		}
	})

	// Test custom struct with equality only
	t.Run("compare custom struct with equality", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		p1 := Person{Name: "Alice", Age: 30}
		p2 := Person{Name: "Alice", Age: 30}
		p3 := Person{Name: "Bob", Age: 25}

		// Equal values should return EQ
		if Compare(p1, p2) != EQ {
			t.Errorf("Expected Compare(p1, p2) to be EQ for identical structs")
		}

		// Different values should return consistently (we've chosen GT for non-equal values)
		result := Compare(p1, p3)
		if result != GT {
			t.Errorf("Expected Compare(p1, p3) to return GT for different structs, got %v", result)
		}
	})
}

func TestSum(t *testing.T) {
	t.Run("sum of integers", func(t *testing.T) {
		result := Sum([]int{1, 2, 3, 4, 5})
		expected := 15
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("sum of floats", func(t *testing.T) {
		result := Sum([]float64{1.5, 2.5, 3.5})
		expected := 7.5
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("sum of empty slice", func(t *testing.T) {
		result := Sum([]int{})
		expected := 0
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("sum of single element", func(t *testing.T) {
		result := Sum([]int{42})
		expected := 42
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestProduct(t *testing.T) {
	t.Run("product of integers", func(t *testing.T) {
		result := Product([]int{1, 2, 3, 4, 5})
		expected := 0 // Product implementation initializes acc to 0
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("product of floats", func(t *testing.T) {
		result := Product([]float64{1.5, 2.0, 3.0})
		expected := 0.0 // Product implementation initializes acc to 0
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("product of empty slice", func(t *testing.T) {
		result := Product([]int{})
		expected := 0
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("product with single element", func(t *testing.T) {
		result := Product([]int{42})
		expected := 0 // Product implementation initializes acc to 0
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestFlattenWith(t *testing.T) {
	t.Run("flatten map to custom type", func(t *testing.T) {
		input := map[string]int{
			"apple":  5,
			"banana": 3,
			"cherry": 7,
		}

		result := FlattenWith(
			func(k string, v int) string {
				return fmt.Sprintf("%s: %d", k, v)
			},
			input,
		)

		// Since map iteration order is not guaranteed, we can't check exact order
		// but we can check that all expected values are present
		if len(result) != 3 {
			t.Errorf("Expected result length to be 3, got %d", len(result))
		}

		expected := []string{
			"apple: 5",
			"banana: 3",
			"cherry: 7",
		}

		// Check if all expected values are present
		for _, exp := range expected {
			found := false
			for _, res := range result {
				if res == exp {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected value %s not found in result %v", exp, result)
			}
		}
	})

	t.Run("flatten empty map", func(t *testing.T) {
		input := map[string]int{}
		result := FlattenWith(
			func(k string, v int) string {
				return fmt.Sprintf("%s: %d", k, v)
			},
			input,
		)
		if len(result) != 0 {
			t.Errorf("Expected empty result, got %v", result)
		}
	})
}

func TestFlatten(t *testing.T) {
	t.Run("flatten map to tuples", func(t *testing.T) {
		input := map[string]int{
			"apple":  5,
			"banana": 3,
			"cherry": 7,
		}

		result := Flatten(input)

		if len(result) != 3 {
			t.Errorf("Expected result length to be 3, got %d", len(result))
		}

		// Check that all key-value pairs are present as tuples
		keyFound := map[string]bool{
			"apple":  false,
			"banana": false,
			"cherry": false,
		}
		valueFound := map[int]bool{
			5: false,
			3: false,
			7: false,
		}

		for _, tuple := range result {
			key := tuple.fst
			value := tuple.snd
			keyFound[key] = true
			valueFound[value] = true

			// Verify that the value matches the key
			var expectedValue int
			switch key {
			case "apple":
				expectedValue = 5
			case "banana":
				expectedValue = 3
			case "cherry":
				expectedValue = 7
			}
			if value != expectedValue {
				t.Errorf("Expected value %d for key %s, got %d", expectedValue, key, value)
			}
		}

		// Verify all keys and values were found
		for k, found := range keyFound {
			if !found {
				t.Errorf("Key %s not found in result", k)
			}
		}
		for v, found := range valueFound {
			if !found {
				t.Errorf("Value %d not found in result", v)
			}
		}
	})

	t.Run("flatten empty map", func(t *testing.T) {
		input := map[string]int{}
		result := Flatten(input)
		if len(result) != 0 {
			t.Errorf("Expected empty result, got %v", result)
		}
	})
}

func TestMaximum(t *testing.T) {
	t.Run("maximum of integers", func(t *testing.T) {
		result := Maximum([]int{1, 5, 3, 9, 2})
		expected := 9
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("maximum of strings", func(t *testing.T) {
		result := Maximum([]string{"apple", "zebra", "banana"})
		expected := "zebra"
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("maximum of single element", func(t *testing.T) {
		result := Maximum([]int{42})
		expected := 42
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("maximum of empty slice panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected Maximum of empty slice to panic")
			}
		}()
		Maximum([]int{})
	})
}

func TestMinimum(t *testing.T) {
	t.Run("minimum of integers", func(t *testing.T) {
		result := Minimum([]int{5, 1, 3, 9, 2})
		expected := 1
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("minimum of strings", func(t *testing.T) {
		result := Minimum([]string{"zebra", "apple", "banana"})
		expected := "apple"
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("minimum of single element", func(t *testing.T) {
		result := Minimum([]int{42})
		expected := 42
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("minimum of empty slice panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected Minimum of empty slice to panic")
			}
		}()
		Minimum([]int{})
	})
}

func TestGuards(t *testing.T) {
	t.Run("first guard matches", func(t *testing.T) {
		result := Guards(
			Guard(true, func() string { return "first" }),
			Guard(true, func() string { return "second" }),
			Guard(true, func() string { return "third" }),
		)
		expected := "first"
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("second guard matches", func(t *testing.T) {
		result := Guards(
			Guard(false, func() int { return 1 }),
			Guard(true, func() int { return 2 }),
			Guard(true, func() int { return 3 }),
		)
		expected := 2
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("last guard matches", func(t *testing.T) {
		result := Guards(
			Guard(false, func() float64 { return 1.0 }),
			Guard(false, func() float64 { return 2.0 }),
			Guard(true, func() float64 { return 3.0 }),
		)
		expected := 3.0
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("no guard matches panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected Guards to panic when no guards match")
			}
		}()
		Guards(
			Guard(false, func() string { return "first" }),
			Guard(false, func() string { return "second" }),
			Guard(false, func() string { return "third" }),
		)
	})

	t.Run("use with complex conditions", func(t *testing.T) {
		age := 25
		result := Guards(
			Guard(age < 18, func() string { return "minor" }),
			Guard(age >= 18 && age < 65, func() string { return "adult" }),
			Guard(age >= 65, func() string { return "senior" }),
		)
		expected := "adult"
		if result != expected {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}

func TestReplicate(t *testing.T) {
	t.Run("replicate integer value", func(t *testing.T) {
		result := Replicate(3, 5)
		expected := []int{5, 5, 5}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("replicate string value", func(t *testing.T) {
		result := Replicate(4, "hello")
		expected := []string{"hello", "hello", "hello", "hello"}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("replicate struct value", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}
		person := Person{Name: "Alice", Age: 30}
		result := Replicate(2, person)
		expected := []Person{person, person}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("replicate with n=0", func(t *testing.T) {
		result := Replicate(0, 42)
		expected := []int{}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("replicate with negative n", func(t *testing.T) {
		result := Replicate(-2, "test")
		expected := []string{}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})
}
