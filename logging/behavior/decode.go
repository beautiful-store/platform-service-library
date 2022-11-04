package behavior

func DecodeMessage(message interface{}) *Log {
	c := message.(logContext)

	return &Log{Context: &c}
}
