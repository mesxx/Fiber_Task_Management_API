package models

type (
	RequestCreateUser struct {
		Name     string `json:"name" validate:"required,min=5,max=20"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=5"`
	}

	RequestLoginUser struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=5"`
	}

	RequestCreateTask struct {
		UserID      uint   `json:"user_id" validate:"required"`
		Title       string `json:"title" validate:"required"`
		Description string `json:"description"`
		Image       string `json:"image"`
	}

	RequestUpdateTask struct {
		UserID      uint    `json:"user_id" validate:"required"`
		Status      string  `json:"status"`
		Title       string  `json:"title" validate:"required"`
		Description *string `json:"description"`
	}
)
