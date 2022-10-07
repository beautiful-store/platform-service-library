package library

import (
	"testing"
)

func TestStruct2Json(t *testing.T) {
	v1 := &struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}{
		ID:   1,
		Name: "홍길동",
	}

	t.Log("origin : ", v1)

	result, err := Struct2Json(v1)
	if err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"id":1,"name":"홍길동"}`
	if result != expected {
		t.Errorf("expected : %s but get : %s", expected, result)
	}
}

// func Struct2Json1(s interface{}) {
// 	Struct2Json(s)
// }

// func Struct2Json2(s interface{}) {
// 	f := func(v interface{}) (string, error) {
// 		if v == nil {
// 			return "", nil
// 		} else {
// 			if e, err := easyjson.Marshal(v.(easyjson.Marshaler)); err != nil {
// 				return "", err
// 			} else {
// 				return string(e), nil
// 			}
// 		}
// 	}

// 	f(s)
// }

// func BenchmarkStruct2Json1(b *testing.B) {
// 	v1 := &(struct {
// 		ID   int    `json:"id"`
// 		Name string `json:"name"`
// 	}{
// 		ID:   1,
// 		Name: "홍길동",
// 	})

// 	for i := 0; i < b.N; i++ {
// 		Struct2Json1(v1)
// 	}
// }

// func BenchmarkStruct2Json2(b *testing.B) {
// 	v1 := &(struct {
// 		ID   int    `json:"id"`
// 		Name string `json:"name"`
// 	}{
// 		ID:   1,
// 		Name: "홍길동",
// 	})

// 	for i := 0; i < b.N; i++ {
// 		Struct2Json1(v1)
// 	}
// }

// go test -bench=.
// go test -bench=. -benchmem
