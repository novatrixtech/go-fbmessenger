package fbmodelrecieve

/*
FacebookMessageRecieved - Facebook Message Received Object
*/
type FacebookMessageRecieved struct {
	Object string `json:"object"`
	Entry  []struct {
		ID        string `json:"id"`
		Time      int64  `json:"time"`
		Messaging []struct {
			Sender struct {
				ID string `json:"id"`
			} `json:"sender"`
			Recipient struct {
				ID string `json:"id"`
			} `json:"recipient"`
			Referral struct {
				Value  string `json:"ref"`
				Source string `json:"source"`
				Type   string `json:"type"`
			} `json:"referral"`
			//Testar colocar todos os campos e ver se o Macaron faz o bind e deixa nulo
			//quando nao tiver esse dado
			Timestamp int64 `json:"timestamp"`
			Read      struct {
				Watermark int `json:"watermark"`
				Seq       int `json:"seq"`
			} `json:"read"`
			Delivery struct {
				Mids []string `json:"mids"`
			} `json:"delivery"`
			Message struct {
				Mid         string `json:"mid"`
				Seq         int    `json:"seq"`
				Text        string `json:"text"`
				IsEcho      bool   `json:"is_echo"`
				AppID       int    `json:"app_id"`
				Attachments []struct {
					Type    string `json:"type"`
					Title   string `json:"title"`
					URL     string `json:"url"`
					Payload struct {
						URL         string `json:"url"`
						Coordinates struct {
							Latitude  float32 `json:"lat"`
							Longitude float32 `json:"long"`
						} `json:"coordinates"`
					} `json:"payload"`
				} `json:"attachments"`
			} `json:"message"`
			Postback struct {
				Payload  string `json:"payload"`
				Referral struct {
					Value  string `json:"ref"`
					Source string `json:"source"`
					Type   string `json:"type"`
				} `json:"referral"`
			} `json:"postback"`
		} `json:"messaging"`
	} `json:"entry"`
}

/*
		Example of a FB JSON Payload recieved:
		{
"object":"page",
"entry":[
		{
				"id":"1070203333093348",
				"time":1478076699002,
				"messaging":[
						{
								"sender":{
										"id":"1070203333093348"
								},
								"recipient":{
										"id":"1160103300748406"
								},
								"timestamp":1478076694758,
								"message":{
										"is_echo":true,
										"app_id":1793577977581304,
										"mid":"mid.1478076694758:402219fe95",
										"seq":197,
										"text":"Jefferson, voc\u00ea escreveu: Teste 37"
								}
						},
						{
								"sender":{
										"id":"1160103300748406"
								},
								"recipient":{
										"id":"1070203333093348"
								},
								"timestamp":0,
								"delivery":{
										"mids":[
												"mid.1478076694758:402219fe95"
										],
										"watermark":1478076694758,
										"seq":198
								}
						},
						{
								"sender":{
										"id":"1160103300748406"
								},
								"recipient":{
										"id":"1070203333093348"
								},
								"timestamp":1478076695056,
								"read":{
										"watermark":1478076694758,
										"seq":199
								}
						},
						{
								"sender":{
										"id":"1070203333093348"
								},
								"recipient":{
										"id":"1160103300748406"
								},
								"timestamp":1478076698947,
								"message":{
										"is_echo":true,
										"app_id":1793577977581304,
										"mid":"mid.1478076698947:56234ec930",
										"seq":200,
										"attachments":[
												{
														"type":"image",
														"payload":{
																"url":"https:\/\/scontent.xx.fbcdn.net\/v\/t34.1-12\/19962480_1080672528713095_1142210083_n.gif?_nc_ad=z-m&oh=332ce5c573defb1de910701b3860bb5b&oe=581B4472"
														}
												}
										]
								}
						}
				]
		}
]
}
*/
