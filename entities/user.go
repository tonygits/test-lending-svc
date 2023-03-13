package entities

type User struct {
	Username  string  `json:"username"`
	UserLimit float64 `json:"user_limit"`
}
