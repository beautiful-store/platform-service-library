package apicall

func DecodeMessage(message interface{}) *APICall {
	c := message.(APICall)

	return &c
}
