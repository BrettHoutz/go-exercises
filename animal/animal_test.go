package animal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpeak(t *testing.T) {
	testCat := &cat{"Turtles", "Tortoiseshell", 1}
	testDog := &dog{"Sasha", "Golden Retriever", 1}
	testBird := &bird{"Mr. Stubs", "Parakeet", 1}

	assert.Equal(t, "Turtles says Meow!", Speak(testCat), "works on a cat")
	assert.Equal(t, "Sasha the Golden Retriever says Bark!", Speak(testDog), "works on a dog")
	assert.Equal(t, "Mr. Stubs the Parakeet says Tweet!", Speak(testBird), "works on a bird")
}
