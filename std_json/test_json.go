package std_json

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

func TestJson() {
	structToJson()
	jsonToStruct()
	testNest()
	structToXml()
}

type DetailStruct struct {
	Hoby string `json:"hoby"`
}
type Xx struct {
	Name   string       `json:"name"`
	Age    int          `json:"age"`
	Job    string       `json:"job"`
	Detail DetailStruct `json:"detail"`
}

func structToJson() {
	s := Xx{
		Name: "xxx",
		Age:  12,
		Job:  "job",
		Detail: DetailStruct{
			Hoby: "hoby",
		},
	}
	fmt.Printf("s: %v\n", s)
	b, _ := json.Marshal(s)
	fmt.Printf("b: %v\n", string(b))
}

func jsonToStruct() {
	s := `{"name":"xxx","age":12,"job":"job","detail":{"hoby":"hoby"}}`
	var target Xx
	json.Unmarshal([]byte(s), &target)
	fmt.Printf("target: %v\n", target)
	fmt.Printf("target: %T\n", target)
	fmt.Printf("target.Detail.Hoby: %v\n", target.Detail.Hoby)
}

func testNest() {
	s := `{"list": ["first","second"]}`
	// go 中 任意类型是 interface{}
	var target interface{}
	json.Unmarshal([]byte(s), &target)
	fmt.Printf("target.(map[string]interface{}): %v\n", target.(map[string]interface{}))
	fmt.Printf("target: %v\n", target)
	// 使用断言进行类型转化
	for _, v := range target.(map[string]interface{}) {
		fmt.Printf("v: %v\n", v)
		for _, x := range v.([]interface{}) {
			fmt.Printf("x: %v\n", x)
		}
	}
}

func structToXml() {
	type DetailStructxml struct {
		Hoby string `xml:"hoby"`
	}

	type Xxxml struct {
		// 最外层包裹改别称
		XMLName xml.Name        `xml:"demo"`
		Name    string          `xml:"name"`
		Age     int             `xml:"age"`
		Job     string          `xml:"job"`
		Detail  DetailStructxml `xml:"detail"`
	}
	s := Xxxml{
		Name: "xxx",
		Age:  12,
		Job:  "job",
		Detail: DetailStructxml{
			Hoby: "hoby",
		},
	}
	fmt.Printf("s: %v\n", s)
	b, _ := xml.Marshal(s)
	fmt.Printf("b: %v\n", string(b))
	b2, _ := xml.MarshalIndent(s, "", "  ")
	fmt.Printf("b2: %v\n", string(b2))
}
