package roc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIdGeneratorGenerate(t *testing.T) {
	generator := &IdGenerator{id: 0}
	assert.Equal(t, 1, generator.Generate())
	ch := make(chan uint32, 100)
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				ch <- generator.Generate()
			}
		}()
	}

	var id uint32
	m := make(map[uint32]bool)
	for i := 0; i < 100; i++ {
		id = <-ch
		m[id] = true
	}

	assert.Equal(t, len(m), 100)
}
