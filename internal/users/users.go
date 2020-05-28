package users

// User is a struct that represent the users we get from database
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
