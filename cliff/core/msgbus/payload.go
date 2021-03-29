package msgbus

const (
	EventVideoReadyForProcessing = "EventVideoReadyForProcessing"
	// EventVideoProcessed          = "EventVideoProcessed"
	// EventVideoProcessFailed      = "EventVideoProcessFailed"
)

type EventSerialization struct {
	Event string
	Slug  string
}
