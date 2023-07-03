package request_body

type CreateEncryptedVault struct {
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
}

type DeleteEncryptedVault struct {
	DeleteCode string `json:"deleteCode"`
}
