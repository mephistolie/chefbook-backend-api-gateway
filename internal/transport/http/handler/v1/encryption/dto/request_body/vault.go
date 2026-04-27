package request_body

type CreateEncryptedVault struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
	Salt       string `json:"salt"`
}

type DeleteEncryptedVault struct {
	DeleteCode string `json:"deleteCode"`
}
