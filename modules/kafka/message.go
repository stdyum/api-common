package kafka

type Message struct {
	Key     string
	Value   string
	Headers map[string]string
}
