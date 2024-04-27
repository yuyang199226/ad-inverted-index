package index

import (
	"fmt"
	"github.com/RoaringBitmap/roaring"
)

type Field struct {
	None        *roaring.Bitmap
	Int64Map    map[int64]*roaring.Bitmap
	StringMap   map[string]*roaring.Bitmap
	IncludeNone bool
}

func NewField() *Field {
	f := Field{
		None:        roaring.NewBitmap(),
		Int64Map:    make(map[int64]*roaring.Bitmap),
		StringMap:   make(map[string]*roaring.Bitmap),
		IncludeNone: true,
	}
	return &f
}

func (f *Field) Get(val any) *roaring.Bitmap {
	switch val.(type) {
	case int64:
		innerVal := val.(int64)
		return f.getInt64(innerVal)
	case string:
		innerVal := val.(string)
		return f.getString(innerVal)
	case []string:
		f.getArrayString(val.([]string))

	case []int64:
		f.getArrayInt64(val.([]int64))
	case nil:
		return f.None
	default:
		panic("unknown type")
	}
	return nil
}

func (f *Field) Add(val any, adID uint32) error {
	switch val.(type) {
	case int64:
		innerVal := val.(int64)
		f.addSingleInt64(innerVal, adID)
	case string:
		innerVal := val.(string)
		f.addSingleString(innerVal, adID)
	case []string:
		f.addArrayString(val.([]string), adID)

	case []int64:
		f.addArrayInt64(val.([]int64), adID)
	case nil:
		f.addNone(adID)
	default:
		panic("unknown type")
	}
	return nil
}

func (f *Field) addSingleInt64(val int64, adID uint32) {
	if f.Int64Map[val] == nil {
		f.Int64Map[val] = roaring.NewBitmap()
	}
	f.Int64Map[val].Add(adID)
}

func (f *Field) addSingleString(val string, adID uint32) {
	if f.StringMap[val] == nil {
		f.StringMap[val] = roaring.NewBitmap()
	}
	f.StringMap[val].Add(adID)
}

func (f *Field) addArrayString(nums []string, adID uint32) {
	for _, v := range nums {
		f.addSingleString(v, adID)
	}
}

func (f *Field) addArrayInt64(nums []int64, adID uint32) {
	for _, v := range nums {
		f.addSingleInt64(v, adID)
	}
}

func (f *Field) addNone(adID uint32) {
	if !f.None.Contains(adID) {
		f.None.Add(adID)
	}
}


func (f *Field) getInt64(val int64) *roaring.Bitmap {
	none := f.None
	valbm := f.Int64Map[val]
	if valbm == nil {
		return none
	}
	return BitmapOr([]*roaring.Bitmap{none, valbm})
} 

func (f *Field) getArrayInt64(nums []int64) *roaring.Bitmap {
	none := f.None
	ls := make([]*roaring.Bitmap, 0)
	ls = append(ls, none)
	for _, val := range nums {
		if f.Int64Map[val] != nil && !f.Int64Map[val].IsEmpty(){
			ls = append(ls, f.Int64Map[val])
		}
	}
	return BitmapOr(ls)
} 

func (f *Field) getArrayString(nums []string) *roaring.Bitmap {
	none := f.None
	ls := make([]*roaring.Bitmap, 0)
	ls = append(ls, none)
	for _, val := range nums {
		if f.StringMap[val] != nil && !f.StringMap[val].IsEmpty(){
			ls = append(ls, f.StringMap[val])
		}
	}
	return BitmapOr(ls)
} 


func (f *Field) getString(val string) *roaring.Bitmap {
	none := f.None
	valbm := f.StringMap[val]
	if valbm == nil {
		return none
	}
	return BitmapOr([]*roaring.Bitmap{none, valbm})
} 



func BitmapOr(bitMaps []*roaring.Bitmap) *roaring.Bitmap {
	switch len(bitMaps) {
	case 0:
		return nil
	case 1:
		return bitMaps[0]
	case 2:
		return roaring.Or(bitMaps[0], bitMaps[1])
	}
	return roaring.HeapOr(bitMaps...)
}


func (f *Field) String() string {
	res := "none: "
	res += fmt.Sprintf("%v\n", *f.None)
	if len(f.Int64Map) != 0 {
		res += "int64: "
		for k, v := range f.Int64Map {
			res += fmt.Sprintf("%d: vv: %v\t", k, v)
		}
	} else {
		for k, v := range f.StringMap {
			res += fmt.Sprintf("%s: vv: %v\t", k, v)
		}
	}
	res += "\n"

	return res
}
