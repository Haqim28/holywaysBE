package users

//
type CreateUserRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Image    string `json:"image" form:"image" `
}

type UpdateUserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password" `
	FullName string `json:"fullname" form:"fullname"`
	Phone    string `json:"phone" form:"phone" `
	Image    string `json:"image" form:"image" `
}
