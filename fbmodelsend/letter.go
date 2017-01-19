package fbmodelsend

/*
Letter is a complete message to a Facebook user.
We use this name to refer a old letter because your message to be delivered
it needed a sender, a reciever not only text in order to the mail company be
able to find the reciever (or recipient).
In this case our mail company is Facebook
*/
//TODO: Voltar Message para ponteiros
type Letter struct {
	Recipient Recipient `json:"recipient"`
	Message   Message   `json:"message"`
}
