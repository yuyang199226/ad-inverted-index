package index

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForwardIndex(t *testing.T) {

	fidx := NewForwardIndex()
	fidx.Set(&Ad{AdID: 1, AdZone: 10})
	ad := fidx.Get(1)
	assert.Equal(t, int32(10), ad.AdZone)

}
