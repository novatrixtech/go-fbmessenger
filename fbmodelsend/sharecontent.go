package fbmodelsend

/*
ShareContents represents a Rich Media content to be shared
*/
type ShareContents struct {
	Attachment struct {
		Type    string `json:"type"`
		Payload struct {
			TemplateType string             `json:"template_type"`
			Elements     *[]TemplateElement `json:"elements"`
		} `json:"payload"`
	} `json:"attachment"`
}
