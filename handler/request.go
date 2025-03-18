package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    int64  `json:"phone"`
	Birth    string `json:"birth"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}
	if r.Phone < 0 {
		return errParamIsRequired("phone", "int64")
	}
	if r.Birth == "" {
		return errParamIsRequired("birth", "string")
	}
	return nil
}

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    int64  `json:"phone"`
	Birth    string `json:"birth"`
}

func (r *UpdateUserRequest) Validate() error {
	fmt.Println(r.Phone)
	if r.Name != "" || r.Email != "" || r.Password != "" || r.Birth != "" || r.Phone > 0 {
		return nil
	}
	return fmt.Errorf("at least one valid filed must be provided")
}
