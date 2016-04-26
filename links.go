package pivnet

type Links struct {
	EULA           map[string]string `json:"eula,omitempty"`
	Download       map[string]string `json:"download,omitempty"`
	ProductFiles   map[string]string `json:"product_files,omitempty"`
	EULAAcceptance map[string]string `json:"eula_acceptance,omitempty"`
}
