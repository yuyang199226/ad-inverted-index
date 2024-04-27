package index

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvertedIndex_Get(t *testing.T) {
	vidx := NewInvertedIndex()
	ad1 := &Ad{
		AdID:   1,
		AdZone: 10,
		Target: &Target{
			Location: []string{"beijing", "shanghai"},
			//Gender:   1,
			Platform: "ios",
		},
	}
	ad2 := &Ad{
		AdID:   10,
		AdZone: 9,
		Target: &Target{

			Network:  []string{"wifi", "4g"},
			Keywords: []string{"music", "happy"},
		},
	}
	ad3 := &Ad{
		AdID:   11,
		AdZone: 9,
		Target: &Target{},
	}
	vidx.Add(ad1)
	vidx.Add(ad2)
	vidx.Add(ad3)
	fmt.Println(vidx)
	res := vidx.Search(&RequestTarget{
		Location: "beijing",
		Network:  "wifi",
		Platform: "ios",
	})
	fmt.Println(res)
	assert.Equal(t, 3, len(res))
}
