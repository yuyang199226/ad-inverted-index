package index

import "reflect"

type Dimension struct {
	Name        string
	Type        string
	IncludeNone bool
}

var DefaultDimensions = []Dimension{
	Dimension{Name: "Location", IncludeNone: true},
	Dimension{Name: "Network", IncludeNone: true},
	Dimension{Name: "Platform", IncludeNone: true},
	Dimension{Name: "Keywords", IncludeNone: true},
	Dimension{Name: "Gender", IncludeNone: true},
}

type Target struct {
	Location []string `json:"location"` // 城市
	Network  []string `json:"network"`  // wifi 4g
	Platform string   `json:"platform"` // ios, android
	Keywords []string `json:"keywords"`
	Gender   int64    `json:"gender"`
}

type RequestTarget struct {
	Location string   `json:"location"` // 城市
	Network  string   `json:"network"`  // wifi 4g
	Platform string   `json:"platform"` // ios, android
	Keywords []string `json:"keywords"`
	Gender   int64     `json:"gender"`
}

func (t *RequestTarget) Get(name string) (any, string) {

	p := reflect.TypeOf(t).Elem()

	ageField, _ := p.FieldByName(name)

	// 获取字段的类型
	fieldType := ageField.Type

	// 使用反射获取结构体字段的值
	v := reflect.ValueOf(t).Elem()
	fieldValue := v.FieldByName(name).Interface()
	//fmt.Println("Field Value:", fieldValue)
	return fieldValue, fieldType.Name()

}

func (t *Target) Get(name string) any {

	s := reflect.ValueOf(*t)
	return s.FieldByName(name)

}
