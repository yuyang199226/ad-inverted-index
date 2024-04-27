package index

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/RoaringBitmap/roaring"
)

const NoneName = "none"

type InvertedIndex struct {
	inverted map[Dimension]*Field
}

func NewInvertedIndex() *InvertedIndex {
	vidx := InvertedIndex{
		inverted: make(map[Dimension]*Field),
	}
	for _, v := range DefaultDimensions {
		vidx.inverted[v] = NewField()
	}
	return &vidx
}
func (vidx *InvertedIndex) Search(target *RequestTarget) []uint32 {
	//res := roaring.New()
	list := make([]*roaring.Bitmap, 0)
	for dimension, field := range vidx.inverted {
		//dimension := k.Name // location
		rValue := reflect.ValueOf(target).Elem().
		FieldByName(dimension.Name)
		val := rValue.Interface()
		dimenRes := field.Get(val)
		if dimenRes != nil {
			list = append(list, dimenRes)
		}
	}

	fmt.Println(len(list), list)

	return And(list).ToArray()
}

func And(list []*roaring.Bitmap) *roaring.Bitmap {
	if len(list) == 0 {
		return roaring.NewBitmap()
	}
	if len(list) == 1 {
		return list[0]
	}
	res := list[0]
	for i := 1; i < len(list); i++ {
		res.And(list[i])
	}
	return res
}

func (vidx *InvertedIndex) Add(ad *Ad) error {
	target := ad.Target
	if target == nil {
		return nil
	}
	for _, dimension := range DefaultDimensions {
		name := dimension.Name // location
		rValue := reflect.ValueOf(target).Elem().
			FieldByName(name)
		val := rValue.Interface()
		if !rValue.IsZero() {
			vidx.inverted[dimension].Add(val, ad.AdID)
		}else {
			vidx.inverted[dimension].Add(nil, ad.AdID)
		}
	}
	return nil
}

func (vidx *InvertedIndex) String() string {

	fmt.Println(len(vidx.inverted))
	res := ""
	for k, v := range vidx.inverted {

		res += fmt.Sprintf("dimension: %s: { %v\t }", k.Name, *v)
		res += "\n"
	}
	return res
}

func isIterable(v string) bool {
	if strings.HasPrefix(v, "[]") {
		return true
	}
	return false
}
