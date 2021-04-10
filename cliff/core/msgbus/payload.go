package msgbus

const (
	EventVideoReadyForProcessing = "EventVideoReadyForProcessing"
	// EventVideoProcessed          = "EventVideoProcessed"
	// EventVideoProcessFailed      = "EventVideoProcessFailed"
)

type EventSerialization struct {
	Event    string `json:"event"`
	Slug     string `json:"slug"`
	Filename string `json:"filename"`
}
