package accumulator

import "testing"

// dedupeSwapDirtResult is a struct that represents the input and expected
// output for dedupeSwapDirt().
type dedupeSwapDirtResult struct {
	a        []uint64
	b        []arrow
	expected []uint64
}

// dedupeTests holds the test cases
var dedupeTests = []dedupeSwapDirtResult{
	dedupeSwapDirtResult{
		a:        []uint64{},
		b:        []arrow{},
		expected: []uint64{},
	},
	dedupeSwapDirtResult{
		a:        []uint64{928, 932, 927, 928, 929, 932},
		b:        []arrow{},
		expected: []uint64{928, 932, 927, 928, 929, 932},
	},
	dedupeSwapDirtResult{
		a: []uint64{},
		b: []arrow{
			arrow{
				from: 256,
				to:   252,
			},
			arrow{
				from: 261,
				to:   258,
			},
			arrow{
				from: 288,
				to:   266,
			},
			arrow{
				from: 298,
				to:   292,
			},
		},
		expected: []uint64{},
	},
	dedupeSwapDirtResult{
		a: []uint64{638, 641, 645, 658},
		b: []arrow{
			arrow{
				from: 643,
				to:   640,
			},
			arrow{
				from: 657,
				to:   656,
			},
		},
		expected: []uint64{638, 641, 645, 658},
	},
	dedupeSwapDirtResult{
		a: []uint64{0, 3, 5, 7},
		b: []arrow{
			arrow{
				from: 2,
				to:   0,
			},
			arrow{
				from: 10,
				to:   9,
			},
		},
		expected: []uint64{3, 5, 7},
	},
	dedupeSwapDirtResult{
		a: []uint64{1, 4, 5, 10, 16},
		b: []arrow{
			arrow{
				from: 1,
				to:   4,
			},
			arrow{
				from: 1,
				to:   11,
			},
			arrow{
				from: 1,
				to:   16,
			},
		},
		expected: []uint64{1, 5, 10},
	},
}

func equalUint64(a []uint64, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}
	for idx, val := range b {
		if a[idx] != val {
			return false
		}
	}
	return true
}

// TestDedupeSwapDirt tests that dedupeSwapDirt() dedupes and modifies the input
// array as described.
func TestDedupeSwapDirt(t *testing.T) {
	for _, test := range dedupeTests {
		result := dedupeSwapDirt(test.a, test.b)
		if !equalUint64(result, test.expected) {
			t.Errorf("Expected result for dedupeSwapDirt: \n%v\nActual Result: \n%v\n", test.expected, result)
			return
		}

		// make sure that the input array _was_ modified, and equals the result,
		// up to the proper length
		if !equalUint64(result, test.a[:len(result)]) {
			t.Errorf("The arguments for dedupeSwapDirt: \n%v\nActual Result: \n%v\n", test.a, result)
			return
		}
	}
	return
}
