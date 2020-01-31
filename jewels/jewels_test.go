package jewels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumJewels(t *testing.T) {
	assert.Equal(t, 3, NumJewels("AB", "xyBzzABy"), "counts total jewels in stones")
	assert.Equal(t, 3, NumJewels("ABB", "xyBzzABy"), "counts total jewels when jewels are not unique")
	assert.Equal(t, 3, NumJewels("AB", "ABA"), "counts total jewels when all stones are jewels")
	assert.Equal(t, 0, NumJewels("AB", "xyzzy"), "counts total jewels when no stones are jewels")
	assert.Equal(t, 3, NumJewels("💎💍👑", "💣👑🧨💎💎💣💣"), "works with emoji 😉")
}
