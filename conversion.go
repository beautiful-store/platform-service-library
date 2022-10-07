package library

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

func Struct2Json(s interface{}) (string, error) {
	if s == nil {
		return "", nil
	}

	e, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	return string(e), nil
}

func Struct2Byte(s interface{}) ([]byte, error) {
	if s == nil {
		return nil, nil
	}

	e, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func Struct2Map(s interface{}) (map[string]interface{}, error) {
	var data map[string]interface{}

	if s == nil {
		return data, nil
	}

	e, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(e, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func Int64ArrayToString(list []int64, delim string) string {
	var buffer bytes.Buffer

	for i := 0; i < len(list); i++ {
		buffer.WriteString(strconv.FormatInt(list[i], 10))
		if i != len(list)-1 {
			buffer.WriteString(delim)
		}
	}

	return buffer.String()
}

func StringToInt64Array(ids string, delim string) ([]int64, error) {
	if ids == "" {
		return nil, errors.New("can't find the paaramter")
	}
	if delim == "" {
		delim = ","
	}

	var err error

	stringIDs := strings.Split(ids, delim)
	int64IDs := make([]int64, len(stringIDs))
	for i, val := range stringIDs {
		if int64IDs[i], err = strconv.ParseInt(val, 10, 64); err != nil {
			return nil, errors.New("can't find the paaramter")
		}
	}

	return int64IDs, nil
}

func Byte2Struct(b []byte) (*map[string]interface{}, error) {
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return &m, nil
}

func CreatedToMap(val interface{}) (*map[string]interface{}, error) {
	b, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}

	return Byte2Struct(b)
}

func Map2Struct(mapData map[string]interface{}, object interface{}) error {
	jsonStr, err := json.Marshal(mapData)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonStr, object); err != nil {
		return err
	}

	return nil
}

func Map2Byte(mapData map[string]interface{}) ([]byte, error) {
	return json.Marshal(mapData)
}
