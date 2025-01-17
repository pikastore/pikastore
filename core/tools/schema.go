package tools

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Jwt_Token string `json:"token"`
}

// Bucket
type Bucket struct {
	Name           string   `json:"bucket_name"`
	AllowedOrigins []string `json:"allowed_origins" gorm:"serializer:json"`
	AllowedMethods []string `json:"allowed_methods" gorm:"serializer:json"`
	AllowedHeaders []string `json:"allowed_headers" gorm:"serializer:json"`
	ExposedHeaders []string `json:"exposed_headers" gorm:"serializer:json"`
}

type Settings struct {
	ServerName     string `json:"server_name"`
	SetupCompleted bool   `json:"SetupCompleted"`
	URL            string `json:"url"`
	API_Key        string `json:"API_Key"`
}
