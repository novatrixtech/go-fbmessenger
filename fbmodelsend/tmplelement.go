package fbmodelsend

/*
TemplateElement - Elements of Facebook Generic Template Message
*/
type TemplateElement struct {
	Title    string    `json:"title,omitempty"`
	ItemURL  string    `json:"item_url,omitempty"`
	ImageURL string    `json:"image_url,omitempty"`
	Subtitle string    `json:"subtitle,omitempty"`
	Buttons  []*Button `json:"buttons,omitempty"`
}
