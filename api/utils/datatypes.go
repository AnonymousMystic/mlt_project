package utils

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type QueryUserIdResults struct {
	Id     string
	Sessid string
}
