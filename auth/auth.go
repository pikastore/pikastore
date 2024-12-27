package auth

//User schema & whatnot
type User struct {
	Usernmae   string `json:"username"`
	Password   string `json:"password"`
	Auth_Token string `json:"token"`
}

//Self explanitory
type Settings struct {
	ServerName string `json:"server_name"`
	URL        string `json:"url"`
	API_Key    string `json:"API_Key"`
	//more to come
}

func Auth() {
	
}