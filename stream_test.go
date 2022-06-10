package stream

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStream(t *testing.T) {

	t.Run("Filter", func(t *testing.T) {
		t.Parallel()
		t.Run("plays nicely with empty array", func(t *testing.T) {
			t.Parallel()
			stream := OfArray([]int{})
			filtered := Filter(stream, func(v int) bool { return v%2 == 0 })
			assert.NotPanics(t, func() {
				array := ToArray(filtered)
				assert.Empty(t, array)
			})
		})

		t.Run("respects the predicate", func(t *testing.T) {
			t.Parallel()
			stream := OfArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
			filtered := Filter(stream, func(v int) bool { return v%2 == 0 })

			array := ToArray(filtered)

			assert.Equal(t, []int{2, 4, 6, 8}, array)
		})
	})

	t.Run("Map", func(t *testing.T) {
		t.Parallel()
		t.Run("plays nicely with empty array", func(t *testing.T) {
			t.Parallel()
			stream := OfArray([]int{})
			filtered := Map(stream, func(v int) bool { return v%2 == 0 })
			assert.NotPanics(t, func() {
				array := ToArray(filtered)
				assert.Empty(t, array)
			})
		})

		t.Run("respects the predicate", func(t *testing.T) {
			t.Parallel()
			stream := OfArray([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
			filtered := Map(stream, func(v int) bool { return v%2 == 0 })

			array := ToArray(filtered)

			assert.Equal(t, []bool{false, true, false, true, false, true, false, true, false}, array)
		})
	})

	t.Run("Skip", func(t *testing.T) {
		t.Run("plays nicely with empty array", func(t *testing.T) {
			t.Parallel()
			stream := Skip(OfArray([]int{}), 2)
			assert.NotPanics(t, func() {
				array := ToArray(stream)
				assert.Empty(t, array)
			})
		})

		t.Run("skips the right amount", func(t *testing.T) {
			start := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
			stream := Skip(OfArray(start), 2)
			end := ToArray(stream)

			assert.Equal(t, start[2:], end)
		})
	})
}
