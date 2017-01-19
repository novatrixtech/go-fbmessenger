package fbmodelsend

/*
User represents Facebook User's data.
TODO: Prepare system to receive float point timezone
E.g.
{"first_name":"Justin","last_name":"Lima","profile_pic":"https:\/\/scontent.xx.fbcdn.net\/v\/t1.0-1\/100096587_15
3759611640695_1676397034886816931_n.jpg?oh=9d663b2cf2c5c9f988666105502e2124&oe=58E0DAFB","locale":"pt_BR","timezone":2.5,"gender":"male"}
*/
type User struct {
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	ProfilePic string  `json:"profile_pic"`
	Locale     string  `json:"locale"`
	Timezone   float64 `json:"timezone"`
	Gender     string  `json:"gender"`
	ID         string  `json:"id,omitempty"`
}
