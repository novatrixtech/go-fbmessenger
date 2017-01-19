package fbmodelsend

/*
MessagePayload - Represents a payload of Facebook Message
*/
type MessagePayload struct {
	TemplateType string             `json:"template_type,omitempty"`
	Elements     []*TemplateElement `json:"elements,omitempty"`
	URL          string             `json:"url,omitempty"`
	Text         string             `json:"text,omitempty"`
	Buttons      []*Button          `json:"buttons,omitempty"`
}
