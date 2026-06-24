package dto

type CreateUserRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}
