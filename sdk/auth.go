package sdk

// AuthProvider represents an authentification provider
type AuthProvider struct {
	Name           string                 `json:"name"`
	Type           string                 `json:"type"`
	Icon           string                 `json:"icon"`
	RedirectMethod string                 `json:"redirect_method"`
	RedirectURL    string                 `json:"redirect_url"`
	Body           map[string]interface{} `json:"body"`
	ContentType    string                 `json:"content_type"`
}

type AuthRequest struct {
	AuthProvider string `json:"auth_provider"`
	State        string `json:"state"`
}
