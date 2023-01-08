package convertor

import (
	"fmt"
	"strconv"
)

func ExampleToBool() {
	cases := []string{"1", "true", "True", "false", "False", "0", "123", "0.0", "abc"}

	for i := 0; i < len(cases); i++ {
		result, _ := ToBool(cases[i])
		fmt.Println(result)
	}

	// Output:
	// true
	// true
	// true
	// false
	// false
	// false
	// false
	// false
	// false
}

func ExampleToBytes() {
	result1, _ := ToBytes(1)
	result2, _ := ToBytes("abc")
	result3, _ := ToBytes(true)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// [0 0 0 0 0 0 0 1]
	// [97 98 99]
	// [116 114 117 101]
}

func ExampleToChar() {
	result1 := ToChar("")
	result2 := ToChar("abc")
	result3 := ToChar("1 2#3")

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// []
	// [a b c]
	// [1   2 # 3]
}

func ExampleToChannel() {
	ch := ToChannel([]int{1, 2, 3})
	result1 := <-ch
	result2 := <-ch
	result3 := <-ch

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	// Output:
	// 1
	// 2
	// 3
}

func ExampleToString() {
	result1 := ToString("")
	result2 := ToString(nil)
	result3 := ToString(0)
	result4 := ToString(1.23)
	result5 := ToString(true)
	result6 := ToString(false)
	result7 := ToString([]int{1, 2, 3})

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)
	fmt.Println(result7)

	// Output:
	//
	//
	// 0
	// 1.23
	// true
	// false
	// [1,2,3]
}

func ExampleToJson() {

	aMap := map[string]int{"a": 1, "b": 2, "c": 3}
	result1, err := ToJson(aMap)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(result1)

	// Output:
	// {"a":1,"b":2,"c":3}
}

func ExampleToFloat() {
	result1, _ := ToFloat("")
	result2, _ := ToFloat("abc")
	result3, _ := ToFloat("-1")
	result4, _ := ToFloat("-.11")
	result5, _ := ToFloat("1.23e3")
	result6, _ := ToFloat(true)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)
	fmt.Println(result6)

	// Output:
	// 0
	// 0
	// -1
	// -0.11
	// 1230
	// 0
}

func ExampleToInt() {
	result1, _ := ToInt("123")
	result2, _ := ToInt("-123")
	result3, _ := ToInt(float64(12.3))
	result4, _ := ToInt("abc")
	result5, _ := ToInt(true)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
	fmt.Println(result5)

	// Output:
	// 123
	// -123
	// 12
	// 0
	// 0
}

func ExampleToPointer() {
	result := ToPointer(123)
	fmt.Println(*result)

	// Output:
	// 123
}

func ExampleToMap() {
	type Message struct {
		name string
		code int
	}
	messages := []Message{
		{name: "Hello", code: 100},
		{name: "Hi", code: 101},
	}
	result := ToMap(messages, func(msg Message) (int, string) {
		return msg.code, msg.name
	})

	fmt.Println(result)

	// Output:
	// map[100:Hello 101:Hi]
}

func ExampleStructToMap() {
	type People struct {
		Name string `json:"name"`
		age  int
	}
	p := People{
		"test",
		100,
	}
	pm, _ := StructToMap(p)

	fmt.Println(pm)

	// Output:
	// map[name:test]
}

func ExampleMapToSlice() {
	aMap := map[string]int{"a": 1, "b": 2, "c": 3}
	result := MapToSlice(aMap, func(key string, value int) string {
		return key + ":" + strconv.Itoa(value)
	})

	fmt.Println(result) //[]string{"a:1", "c:3", "b:2"} (random order)
}

func ExampleColorHexToRGB() {
	colorHex := "#003366"
	r, g, b := ColorHexToRGB(colorHex)

	fmt.Println(r, g, b)

	// Output:
	// 0 51 102
}

func ExampleColorRGBToHex() {
	r := 0
	g := 51
	b := 102
	colorHex := ColorRGBToHex(r, g, b)

	fmt.Println(colorHex)

	// Output:
	// #003366
}

func ExampleEncodeByte() {
	byteData, _ := EncodeByte("abc")
	fmt.Println(byteData)

	// Output:
	// [6 12 0 3 97 98 99]
}

func ExampleDecodeByte() {
	var obj string
	byteData := []byte{6, 12, 0, 3, 97, 98, 99}
	err := DecodeByte(byteData, &obj)
	if err != nil {
		return
	}

	fmt.Println(obj)

	// Output:
	// abc
}
