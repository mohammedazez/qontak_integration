package vendors

type CredentialObject struct {
	ClientID int    `json:"client_id"`
	Channel  string `json:"channel"`
	Token    string `json:"token"`
}

type Credential struct {
	Channel  string `json:"channel"`
	Token    string `json:"token"`
	AuthType int    `json:"auth_type"`
}
