package index

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestField_Add(t *testing.T) {
	f := NewField()
	f.Add(int64(10), 1000)
	assert.Equal(t, 1, len(f.Int64Map))
	fmt.Println(f)
	f.Add(int64(11), 1000)
	fmt.Println(f)

}
