package library

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

// go test -cover
func TestStruct2Byte(t *testing.T) {
	v1 := &struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   1,
		Name: "홍길동",
	}

	result, err := Struct2Byte(v1)
	if err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"id":1,"name":"홍길동"}`
	if string(result) != expected {
		t.Errorf("expected : %s but get : %s", expected, result)
	}
}

func TestStruct2Map(t *testing.T) {
	v1 := &struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   1,
		Name: "홍길동",
	}

	result, err := Struct2Map(v1)
	if err != nil {
		t.Fatal(err.Error())
	}

	expected := make(map[string]interface{})
	expected["id"] = 1
	expected["name"] = "홍길동"
	if fmt.Sprintf("%v", result["id"]) != fmt.Sprintf("%v", expected["id"]) || result["name"] != expected["name"] {
		t.Errorf("expected : %s but get : %s", expected, result)
	}
}

func TestInt64ArrayToString(t *testing.T) {
	v1 := []int64{1, 2, 3}
	result := Int64ArrayToString(v1, ",")

	expected := "1,2,3"

	if result != expected {
		t.Errorf("expected : %s but get : %s", expected, result)
	}
}

func TestStringToInt64Array(t *testing.T) {
	v1 := "1,2,3"
	result, err := StringToInt64Array(v1, ",")

	expected := []int64{1, 2, 3}
	if err != nil {
		t.Fatal(err.Error())
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected : %v but get : %v", expected, result)
	}
}

func TestByte2Struct(t *testing.T) {
	type s struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	v1 := s{
		ID:   1,
		Name: "홍길동",
	}

	v2 := &s{}

	b, _ := json.Marshal(v1)

	err := Byte2Struct(b, &v2)
	if err != nil {
		t.Fatal(err.Error())
	}

	if v1.Name != v2.Name {
		t.Errorf("expected : %v but get : %v", v1.Name, v2.Name)
	}
}

func TestByte2Map(t *testing.T) {
	v1 := &struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   1,
		Name: "홍길동",
	}

	b, _ := json.Marshal(v1)

	result, err := Byte2Map(b)

	if err != nil {
		t.Fatal(err.Error())
	}

	if v1.Name != (*result)["name"] {
		t.Errorf("expected : %v but get : %v", v1.Name, (*result)["name"])
	}
}

func TestMap2Struct(t *testing.T) {
	v1 := make(map[string]interface{})
	v1["id"] = 1
	v1["name"] = "홍길동"

	s := &(struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{})

	err := Map2Struct(v1, s)
	if err != nil {
		t.Fatal(err.Error())
	}

	if v1["name"] != s.Name {
		t.Errorf("expected : %v but get : %v", v1["name"], s.Name)
	}
}

func TestString2Struct(t *testing.T) {
	str := string(`{"id":1, "name":"홍길동"}`)
	v1 := struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{}

	if err := String2Struct(str, &v1); err != nil {
		t.Errorf(err.Error())
	}
}
