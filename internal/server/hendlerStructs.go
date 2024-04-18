package server

type newUser struct {
	Phone string `json:"phone"`
}

type newExecutor struct {
	Login    string `json:"login"`
	Password []byte `json:"password"`
	Phone    string `json:"phone"`
}
