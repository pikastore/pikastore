package core

type User struct {
	Usernmae   string `json:"username"`
	Password   string `json:"password"`
	Auth_Token string `json:"token"`
}

// Bucket
type Bucket struct {
	Name           string   `json:"bucket_name"`
	AllowedOrigins []string `json:"allowed_origins"`
	AllowedMethods []string `json:"allowed_methods"`
	AllowedHeaders []string `json:"allowed_headers"`
	ExposedHeaders []string `json:"exposed_headers"`
}

type Settings struct {
	ServerName string `json:"server_name"`
	URL        string `json:"url"`
	API_Key    string `json:"API_Key"`
}
