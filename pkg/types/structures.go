package types

import validation "github.com/go-ozzo/ozzo-validation"

// response struct | marshalled into json fromat from struct
type BookRequest struct {
	ID          uint   `json:"bookID"`
	BookName    string `json:"bookName"`
	Author      string `json:"author"`
	Publication string `json:"publication,omitempty"`
}

func (book BookRequest) Validate() error {
	return validation.ValidateStruct(&book,
		validation.Field(&book.BookName,
			validation.Required.Error("Book name cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&book.Author,
			validation.Required.Error("Author name cannot be empty"),
			validation.Length(1, 50)))
}

// Auth request struct
type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (auth AuthRequest) Validate() error {
	return validation.ValidateStruct(&auth,
		validation.Field(&auth.Username,
			validation.Required.Error("Username cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&auth.Password,
			validation.Required.Error("Password cannot be empty"),
			validation.Length(1, 50)))
}

// User request struct
type UserRequest struct {
	ID       uint   `json:"userID"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user UserRequest) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Username,
			validation.Required.Error("Username cannot be empty"),
			validation.Length(1, 50)),
		validation.Field(&user.Password,
			validation.Required.Error("Password cannot be empty"),
			validation.Length(1, 50)))
}
