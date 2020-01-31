package comparetriplets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareTriplets(t *testing.T) {
	aScore, bScore := CompareTriplets([]int{1, 2, 3}, []int{1, 2, 3})
	assert.Equal(t, 0, aScore, "gives aScore as 0 when triplets are identical")
	assert.Equal(t, 0, bScore, "gives bScore as 0 when triplets are identical")

	aScore, bScore = CompareTriplets([]int{1, 2, 4}, []int{1, 3, 3})
	assert.Equal(t, 1, aScore, "gives aScore as 1 when one score is higher")
	assert.Equal(t, 1, bScore, "gives bScore as 1 when one score is higher")

	aScore, bScore = CompareTriplets([]int{2, 2, 4}, []int{1, 3, 3})
	assert.Equal(t, 2, aScore, "gives aScore as 2 when two scores are higher")

	aScore, bScore = CompareTriplets([]int{2, 5, 4}, []int{1, 3, 3})
	assert.Equal(t, 3, aScore, "gives aScore as 3 when three scores are higher")

	aScore, bScore = CompareTriplets([]int{1, 2, 4}, []int{5, 3, 3})
	assert.Equal(t, 2, bScore, "gives bScore as 2 when two scores are higher")
}
