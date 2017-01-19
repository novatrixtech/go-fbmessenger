package fbmodelsend

/*
QuickReply - Represents a Facebook's quick reply items
*/
type QuickReply struct {
	ContentType string `json:"content_type"`
	Title       string `json:"title,omitempty"`
	Payload     string `json:"payload,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
}
