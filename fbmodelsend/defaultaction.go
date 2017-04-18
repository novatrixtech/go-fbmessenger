package fbmodelsend

/*
DefaultAction represents default action to be performed by generic elements in generic templates
*/
type DefaultAction struct {
	Type string `json:"type,omitempty"`
	URL  string `json:"url,omitempty"`
}
