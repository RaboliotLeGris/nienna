package msgbus

const (
	EventVideoReadyForProcessing = "EventVideoReadyForProcessing"
)

type EventSerialization struct {
	Event   string `json:"event"`
	Slug    string `json:"slug"`
	Content string `json:"content"`
}
