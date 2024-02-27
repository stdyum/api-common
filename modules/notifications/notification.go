package notifications

type Notification struct {
	TemplateID string            `json:"templateID,omitempty"`
	Email      string            `json:"email,omitempty"`
	Data       map[string]string `json:"data,omitempty"`
}
