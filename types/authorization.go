package types

type RealmAccess struct {
	Roles []string `json:"roles"`
}

type ResourceAccess struct {
	Account map[string][]string `json:"account"`
}

type TokenPayload struct {
	Exp int `json:"exp"`
	Iat int `json:"iat"`
	Jti string `json:"jti"`
	Iss string `json:"iss"`
	Aud string `json:"aud"`
	Typ string `json:"typ"`
	Azp string `json:"azp"`
	RealmAccess RealmAccess `json:"realm_access"`
	ResourceAccess ResourceAccess `json:"resource_access"`
	Scope string `json:"scope"`
	EmailVerified bool `json:"email_verified"`
}

