package fblib

var fbAccessToken string

/*
DefineAccessToken defines the Access Token to be used by the package to call Fb Messenger Endpoints
*/
func DefineAccessToken(ac string) {
	fbAccessToken = ac
}
