package godash

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	t.Run("integers", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Map(input, func(e int) int {
			return e * 13
		})

		assert.EqualValues(t, []int{13, 26, 39}, result)
	})

	t.Run("different types", func(t *testing.T) {
		var mapper = func(s string) int {
			return len(s)
		}
		result := Map([]string{"alice", "bob", "charlie"}, mapper)
		assert.EqualValues(t, []int{5, 3, 7}, result)
	})

	t.Run("empty", func(t *testing.T) {
		result := Map([]int{}, func(e int) int {
			return e * 13
		})

		assert.Len(t, result, 0)
	})
}

func IsOddInteger(i int) bool {
	return i%2 == 1
}

func TestFilter(t *testing.T) {

	t.Run("integers", func(t *testing.T) {
		input := []int{1, 2, 3}
		result := Filter(input, IsOddInteger)

		assert.EqualValues(t, []int{1, 3}, result)
	})
}

func TestForEach(t *testing.T) {
	t.Run("call every time", func(t *testing.T) {
		output := make([]int, 0, 3)
		// test iterator effectively makes a copy of slice by means of side effects
		var iteratee = func(n int) {
			output = append(output, n)
		}
		input := []int{1, 2, 3}
		ForEach(input, iteratee)

		assert.Equal(t, input, output)
	})
}

func TestContains(t *testing.T) {
	t.Run("find", func(t *testing.T) {
		found := Contains([]string{"alice", "bob", "charlie"}, "bob")
		assert.Truef(t, found, "Element should be in slice")
	})

	t.Run("not found", func(t *testing.T) {
		found := Contains([]string{"alice", "bob", "charlie"}, "bobcat")
		assert.False(t, found, "Element should not ®be in slice")
	})
}
