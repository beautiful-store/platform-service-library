package apicall

import (
	"testing"

	lib "github.com/beautiful-store/platform-service-library"
)

func TestDecodeMessage(t *testing.T) {
	a := APICall{Env: "test", ModuleName: "test", LogType: "test"}

	log := DecodeMessage(a)

	t.Log(lib.Struct2Json(log))
}
